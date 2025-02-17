package validation

import (
    "fmt"
    "github.com/xeipuuv/gojsonschema"
    "github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

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
        logger.Info(fmt.Printf("Loading schema for version %s\n", version))
        if err := definitions.LoadSchema(version, "input"); err != nil {
            logger.Error(err)
            return err
        }
    }
    // Proceed with validation
    errors := validateJSON(jsonString, version, "input")
    if errors != nil {
        logger.Error(errors)
        return errors
    }
    return nil
}

func ValidateOutput(jsonString string, version string) error {
    if _, ok := definitions.SchemaStore[version]["output"]; !ok {
        logger.Info(fmt.Printf("Loading schema for version %s\n", version))
        if err := definitions.LoadSchema(version, "output"); err != nil {
            logger.Error(err)
            return err
        }
    }
    // Proceed with validation
    errors := validateJSON(jsonString, version, "output")
    if errors != nil {
        logger.Error(errors)
        return errors
    }
    return nil
}

func ValidateCost(jsonString string, version string) error {
    if _, ok := definitions.SchemaStore[version]["cost"]; !ok {
        logger.Info(fmt.Printf("Loading schema for version %s\n", version))
        if err := definitions.LoadSchema(version, "cost"); err != nil {
            logger.Error(err)
            return err
        }
    }
    // Proceed with validation
    errors := validateJSON(jsonString, version, "cost")
    if errors != nil {
        logger.Error(errors)
        return errors
    }
    return nil
}