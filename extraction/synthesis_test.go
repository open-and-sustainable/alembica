package extraction

import (
    "errors"
    "testing"
)

func TestBasicErrorReporting(t *testing.T) {
    tests := []struct {
        name        string
        err        error
        expectError bool
    }{
        {
            name: "No error case",
            err:  nil,
            expectError: false,
        },
        {
            name: "Simulated error",
            err:  errors.New("simulated failure"),
            expectError: true,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            if tc.expectError && tc.err == nil {
                t.Errorf("%s: expected error, got nil", tc.name)
            } else if !tc.expectError && tc.err != nil {
                t.Errorf("%s: unexpected error: %v", tc.name, tc.err)
            }
        })
    }
}
