// File: sharedlib/export.go
package main

/*
#include <stdlib.h>
*/
import "C"

import (
    "fmt"
    "github.com/open-and-sustainable/alembica"
    "unsafe"
)

// Common error handling and memory management functions
func handlePanic() *C.char {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
        return C.CString(fmt.Sprint(r))
    }
    return nil
}

// Python-specific function
//export RunSynthesisPython
func RunSynthesisPython(input *C.char) *C.char {
    defer handlePanic()
    goInput := C.GoString(input)
    err := alembica.RunSynthesis(goInput)
    if err != nil {
        return C.CString(err.Error())
    }
    return nil
}

// R-specific function
//export RunSynthesisR
func RunSynthesisR(input *C.char) *C.char {
    defer handlePanic()
    goInput := C.GoString(input)
    err := alembica.RunSynthesis(goInput)
    if err != nil {
        return C.CString(err.Error())
    }
    return C.CString("Review completed successfully")
}

// Free memory function used by both interfaces
//export FreeCString
func FreeCString(str *C.char) {
    C.free(unsafe.Pointer(str))
}

func main() {}