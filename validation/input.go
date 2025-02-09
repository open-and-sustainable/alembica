package validation

import (
    "github.com/open-and-sustainable/alembica/utils/definitions"
)

func ValidateInput(jsonString string, version string) ([]string, error) {
    // Ensure the schema is loaded (this could be skipped if preloading all schemas on init)
    if _, ok := definitions.SchemaStore[version]["input"]; !ok {
        if err := definitions.LoadSchema(version, "input"); err != nil {
            return nil, err
        }
    }

    // Proceed with validation
    return ValidateJSON(jsonString, version, "input")
}