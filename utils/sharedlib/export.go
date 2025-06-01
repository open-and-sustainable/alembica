// File: sharedlib/export.go
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

// Common error handling and memory management functions
func handlePanic() *C.char {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
		return C.CString(fmt.Sprint(r))
	}
	return nil
}

// Python- R- and others-specific function
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

// Free memory function used by both interfaces
//
//export FreeCString
func FreeCString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

func main() {}
