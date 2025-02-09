package definitions

import (
    "embed"
    "fmt"
    "github.com/xeipuuv/gojsonschema"
)

// Embed all schema files in the definitions directory.
//go:embed v*/*
var schemas embed.FS

// SchemaStore holds different versions and types of JSON schemas.
var SchemaStore = make(map[string]map[string]*gojsonschema.Schema)

// LoadSchema loads embedded schema into the store for a given version and type.
func LoadSchema(version, schemaType string) error {
    schemaPath := fmt.Sprintf("definitions/%s/%s_schema.json", version, schemaType)
    schemaData, err := schemas.ReadFile(schemaPath)
    if err != nil {
        return fmt.Errorf("failed to read embedded schema: %v", err)
    }
    loader := gojsonschema.NewBytesLoader(schemaData)
    schema, err := gojsonschema.NewSchema(loader)
    if err != nil {
        return fmt.Errorf("failed to load schema: %v", err)
    }
    if SchemaStore[version] == nil {
        SchemaStore[version] = make(map[string]*gojsonschema.Schema)
    }
    SchemaStore[version][schemaType] = schema
    return nil
}
