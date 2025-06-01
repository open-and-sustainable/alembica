// Package main provides C-compatible shared library exports for alembica functionality.
// It enables non-Go languages like Python, R, and others to use the core features
// of alembica through a C interface.
//
// The exported functions provide a bridge to the main alembica functionalities:
// - Input validation
// - Cost computation
// - Data extraction
//
// Memory management is handled through the FreeCString function, which should be
// called by the client code to free memory allocated by the exported functions.
//
// Example usage from C:
//
//	#include <stdio.h>
//	#include "libsharedlib.h"
//
//	int main() {
//	    const char* input = "{\"metadata\":{...},\"models\":[...],\"prompts\":[...]}";
//	    char* result = RunExtraction(input);
//	    printf("Result: %s\n", result);
//	    FreeCString(result);
//	    return 0;
//	}
package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/open-and-sustainable/alembica/extraction"
	"github.com/open-and-sustainable/alembica/pricing"
	"github.com/open-and-sustainable/alembica/validation"
)

// handlePanic is an internal helper function that recovers from panics in exported functions
// and returns an appropriate error message as a C string.
func handlePanic() *C.char {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
		return C.CString(fmt.Sprint(r))
	}
	return nil
}

// RunValidationInput validates an input JSON string against the specified schema version.
// It returns a success message if validation passes, or an error message if it fails.
//
// Parameters:
//   - input: A C string containing the JSON to validate
//   - version: A C string containing the schema version to validate against
//
// Returns:
//   - A C string containing the validation result or error message.
//     Must be freed with FreeCString() when no longer needed.
//
//export RunValidationInput
func RunValidationInput(input *C.char, version *C.char) *C.char {
	defer handlePanic()
	goInput := C.GoString(input)
	goVersion := C.GoString(version)
	err := validation.ValidateInput(goInput, goVersion)
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString("Input validation successful")
}

// RunComputeCosts calculates the estimated cost of processing the given prompts
// with the specified models. It returns a JSON string with detailed cost information.
//
// Parameters:
//   - input: A C string containing the JSON with prompts and models
//
// Returns:
//   - A C string containing the computed costs as JSON or an error message.
//     Must be freed with FreeCString() when no longer needed.
//
//export RunComputeCosts
func RunComputeCosts(input *C.char) *C.char {
	defer handlePanic()
	goInput := C.GoString(input)

	cost, err := pricing.ComputeCosts(goInput)
	if err != nil {
		// Return a JSON error response instead of a raw error string
		errorResponse := fmt.Sprintf(`{"error": {"message": "%s", "code": 400}}`, err.Error())
		return C.CString(errorResponse)
	}

	return C.CString(cost)
}

// RunExtraction processes the given input to extract structured information using
// the specified LLM models. It handles prompt sequencing, model interaction, and
// result validation.
//
// Parameters:
//   - input: A C string containing the JSON with prompts, models, and extraction parameters
//
// Returns:
//   - A C string containing the extraction results as JSON or an error message.
//     Must be freed with FreeCString() when no longer needed.
//
//export RunExtraction
func RunExtraction(input *C.char) *C.char {
	defer handlePanic()
	goInput := C.GoString(input)

	output, err := extraction.Extract(goInput)
	if err != nil {
		// Return a JSON error response instead of a raw error string
		errorResponse := fmt.Sprintf(`{"error": {"message": "%s", "code": 400}}`, err.Error())
		return C.CString(errorResponse)
	}

	return C.CString(output)
}

// FreeCString frees the memory allocated for a C string that was returned by
// any of the exported functions in this package.
//
// Parameters:
//   - str: A pointer to the C string to free
//
//export FreeCString
func FreeCString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

func main() {}
