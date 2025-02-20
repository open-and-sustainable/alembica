package validation

import (
    "fmt"
    "strings"
    "github.com/xeipuuv/gojsonschema"
    "github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"
)

// validateJSON checks if a given JSON string adheres to the specified schema version and type.
//
// Parameters:
//   - jsonString: The JSON string to be validated.
//   - version: The schema version to validate against.
//   - schemaType: The type of schema (e.g., "input", "output", "cost").
//
// Returns:
//   - An error if validation fails, or nil if validation succeeds.
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

// ValidateInput checks whether the given input JSON conforms to the expected schema.
//
// Parameters:
//   - jsonString: The JSON string to validate.
//   - version: The schema version to use.
//
// Returns:
//   - An error if validation fails, or nil if it succeeds.
func ValidateInput(jsonString string, version string) error {
    if _, ok := definitions.SchemaStore[version]["input"]; !ok {
        logger.Info(fmt.Sprintf("Loading schema for version %s\n", version))
        if err := definitions.LoadSchema(version, "input"); err != nil {
            logger.Error(err)
            return err
        }
    }
    return validateJSON(jsonString, version, "input")
}

// ValidateOutput checks whether the given output JSON conforms to the expected schema.
//
// Parameters:
//   - jsonString: The JSON string to validate.
//   - version: The schema version to use.
//
// Returns:
//   - An error if validation fails, or nil if it succeeds.
func ValidateOutput(jsonString string, version string) error {
    if _, ok := definitions.SchemaStore[version]["output"]; !ok {
        logger.Info(fmt.Sprintf("Loading schema for version %s\n", version))
        if err := definitions.LoadSchema(version, "output"); err != nil {
            logger.Error(err)
            return err
        }
    }
    return validateJSON(jsonString, version, "output")
}

// ValidateCost checks whether the given cost JSON conforms to the expected schema.
//
// Parameters:
//   - jsonString: The JSON string to validate.
//   - version: The schema version to use.
//
// Returns:
//   - An error if validation fails, or nil if it succeeds.
func ValidateCost(jsonString string, version string) error {
    if _, ok := definitions.SchemaStore[version]["cost"]; !ok {
        logger.Info(fmt.Sprintf("Loading schema for version %s\n", version))
        if err := definitions.LoadSchema(version, "cost"); err != nil {
            logger.Error(err)
            return err
        }
    }
    return validateJSON(jsonString, version, "cost")
}
