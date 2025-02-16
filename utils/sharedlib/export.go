// File: sharedlib/export.go
package main

/*
#include <stdlib.h>
*/
import "C"

import (
    "fmt"
    "unsafe"

    "github.com/open-and-sustainable/alembica/validation"
    "github.com/open-and-sustainable/alembica/pricing"
)

// Common error handling and memory management functions
func handlePanic() *C.char {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
        return C.CString(fmt.Sprint(r))
    }
    return nil
}

// Python- and R-specific function
//export RunValidationInput
func RunValidationInput(input *C.char, version *C.char) *C.char {
    defer handlePanic()
    goInput := C.GoString(input)
    goVersion := C.GoString(version)
    err := validation.ValidateInput(goInput, goVersion)
    if err != nil {
        return C.CString(err.Error())
    }
    return C.CString("Input validation successfull")
}
//export RunComputeCosts
func RunComputeCosts(input *C.char) *C.char {
    defer handlePanic()
    goInput := C.GoString(input)
    cost := pricing.ComputeCosts(goInput)
    return C.CString(cost)
}

// Free memory function used by both interfaces
//export FreeCString
func FreeCString(str *C.char) {
    C.free(unsafe.Pointer(str))
}

func main() {}