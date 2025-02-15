package definitions

import (
    "embed"
    "fmt"
    "github.com/xeipuuv/gojsonschema"
    "log"
)

// Embed all schema files in the definitions directory.
//go:embed v*/*
var schemas embed.FS

// SchemaStore holds different versions and types of JSON schemas.
var SchemaStore = make(map[string]map[string]*gojsonschema.Schema)

func LoadSchema(version, schemaType string) error {
    schemaPath := fmt.Sprintf("%s/%s_schema.json", version, schemaType)

    // List all embedded files for debugging
    _, err := schemas.ReadDir(".")
    if err != nil {
        return fmt.Errorf("failed to read embedded files: %v", err)
    }

    schemaData, err := schemas.ReadFile(schemaPath)
    if err != nil {
        return fmt.Errorf("no schema file found for version %s and type %s at %s", version, schemaType, schemaPath)
    }

    loader := gojsonschema.NewBytesLoader(schemaData)
    schema, err := gojsonschema.NewSchema(loader)
    if err != nil {
        return fmt.Errorf("failed to load schema %s: %v", schemaPath, err)
    }

    if SchemaStore[version] == nil {
        SchemaStore[version] = make(map[string]*gojsonschema.Schema)
    }
    SchemaStore[version][schemaType] = schema
    return nil
}


// Initialize SchemaStore at package load.
func init() {
    versions := []string{"v1"} // Add other versions if needed
    schemaTypes := []string{"input", "output"}

    for _, version := range versions {
        for _, schemaType := range schemaTypes {
            if err := LoadSchema(version, schemaType); err != nil {
                log.Printf("Warning: Failed to load schema %s/%s: %v", version, schemaType, err)
            }
        }
    }
}
