package validation

import (
    "fmt"
    "github.com/xeipuuv/gojsonschema"
    "github.com/open-and-sustainable/alembica/utils/definitions"
)

// ValidateJSON validates a JSON string against a specified version and type of schema.
// It returns a slice of error messages if the validation fails, or nil if it succeeds.
func ValidateJSON(jsonString, version, schemaType string) ([]string, error) {
    // Retrieve the schema from the SchemaStore based on the version and type provided.
    schemaMap, versionExists := definitions.SchemaStore[version]
    if !versionExists {
        return nil, fmt.Errorf("no schemas found for version %s", version)
    }

    schema, typeExists := schemaMap[schemaType]
    if !typeExists {
        return nil, fmt.Errorf("no schema found for type %s in version %s", schemaType, version)
    }

    // Load the JSON string into a Loader.
    documentLoader := gojsonschema.NewStringLoader(jsonString)

    // Perform the validation.
    result, err := schema.Validate(documentLoader)
    if err != nil {
        return nil, fmt.Errorf("error during validation: %v", err)
    }

    // Check if there were any validation errors.
    if result.Valid() {
        // If the JSON is valid according to the schema, return nil.
        return nil, nil
    } else {
        // If there are errors, collect all error descriptions to return.
        var errors []string
        for _, desc := range result.Errors() {
            // Append each error message to the slice.
            errors = append(errors, desc.String())
        }
        return errors, nil
    }
}
