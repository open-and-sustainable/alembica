package validation

import (
    "fmt"
    "github.com/xeipuuv/gojsonschema"
    "github.com/open-and-sustainable/alembica/definitions"

    "strings"
)

// ValidateJSON validates a JSON string against a specified version and type of schema.
// It returns an error if validation fails, or nil if it succeeds.
func validateJSON(jsonString, version, schemaType string) error {
    schemaMap, versionExists := definitions.SchemaStore[version]
    if !versionExists {
        return fmt.Errorf("no schemas found for version %s", version)
    }

    schema, typeExists := schemaMap[schemaType]
    if !typeExists {
        return fmt.Errorf("no schema found for type %s in version %s", schemaType, version)
    }

    documentLoader := gojsonschema.NewStringLoader(jsonString)
    result, err := schema.Validate(documentLoader)
    if err != nil {
        return fmt.Errorf("error during validation: %v", err)
    }

    if result.Valid() {
        return nil
    }

    var errorMessages []string
    for _, desc := range result.Errors() {
        errorMessages = append(errorMessages, desc.String())
    }

    return fmt.Errorf("validation errors: %s", strings.Join(errorMessages, "; "))
}

func ValidateInput(jsonString string, version string) error {
    if _, ok := definitions.SchemaStore[version]["input"]; !ok {
        fmt.Printf("Loading schema for version %s\n", version) // Debug log
        if err := definitions.LoadSchema(version, "input"); err != nil {
            return err
        }
    }

    // Proceed with validation
    return validateJSON(jsonString, version, "input")
}
