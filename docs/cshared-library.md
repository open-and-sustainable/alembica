---
title: C-Shared Library
layout: default
---

# C-Shared Library  

## Download  

The shared library can be downloaded as an artifact from the [release page](https://github.com/open-and-sustainable/alembica/releases).  

Once downloaded, make sure it is placed in an accessible path for your system (`/usr/local/lib/`, `C:\path\to\lib\`, etc.).  

## Available Functions  

This library provides three main entry points:  

1. **RunValidationInput** – Validates the input based on a given version.  
2. **RunComputeCosts** – Computes costs based on the provided input.  
3. **RunExtraction** – Extracts data based on the provided input.  

Each function takes a **C string (`char*`)** as input and returns a **C string (`char*`)**, which should be freed using **FreeCString** after use.  

## Language Bindings  

This section provides examples of how to use the **`alembica.so`** shared library in different programming languages. The library exposes three main functions:  

- **`RunValidationInput(input, version)`** → Validates the input based on a given version.  
- **`RunComputeCosts(input)`** → Computes costs based on the provided input.  
- **`RunExtraction(input)`** → Extracts data based on the provided input.  

### How It Works  

Each function:  
- Accepts **a single string or multiple strings as input (`char*`)**.  
- Returns **a dynamically allocated C string (`char*`)**, which must be freed using **`FreeCString`** after use to avoid memory leaks.  

### Usage in Different Languages  

Below are examples of how to call these functions from different languages, using **Foreign Function Interface (FFI)** or the appropriate interop mechanisms.  

- Make sure **`alembica.so`** is in a directory where the runtime can find it (e.g., `/usr/local/lib/` on Linux or in the same directory as your script).  
- Each example demonstrates calling **one of the three functions** and properly freeing the returned string.  


### **1. C**  

```c
#include <stdio.h>  
#include <stdlib.h>  
#include "alembica.h"  

int main() {  
    char *result = RunValidationInput("some_input", "v1.0");  
    printf("Validation Result: %s\n", result);  
    FreeCString(result);  
    return 0;  
}  
```  

### **2. C++**  

```c++
#include <iostream>  
#include "alembica.h"  

int main() {  
    char *result = RunComputeCosts("some_input");  
    std::cout << "Cost Computation: " << result << std::endl;  
    FreeCString(result);  
    return 0;  
}  
```  

### **3. Python (ctypes)**  

```python
import ctypes  

lib = ctypes.CDLL("./alembica.so")  
lib.RunExtraction.restype = ctypes.c_char_p  

result = lib.RunExtraction(b"some_input")  
print("Extraction Result:", result.decode())  
lib.FreeCString(result)  
```  

### **4. Rust (libloading FFI)**  

```rust
use libloading::{Library, Symbol};  
use std::ffi::CString;  

fn main() {  
    let lib = Library::new("alembica.so").unwrap();  
    unsafe {  
        let func: Symbol<unsafe extern "C" fn(*const i8) -> *mut i8> = lib.get(b"RunComputeCosts").unwrap();  
        let input = CString::new("some_input").unwrap();  
        let result = func(input.as_ptr());  
        println!("Result: {}", std::ffi::CStr::from_ptr(result).to_str().unwrap());  
    }  
}  
```  

### **5. Java (JNI)**  

```java
public class SharedLibTest {  
    static { System.loadLibrary("alembica"); }  

    public native String RunExtraction(String input);  

    public static void main(String[] args) {  
        SharedLibTest lib = new SharedLibTest();  
        String result = lib.RunExtraction("some_input");  
        System.out.println("Extraction Result: " + result);  
    }  
}  
```  

### **6. C# (P/Invoke)**  

```c#
using System;  
using System.Runtime.InteropServices;  

class Program {  
    [DllImport("alembica.so")]  
    public static extern IntPtr RunValidationInput(string input, string version);  

    [DllImport("alembica.so")]  
    public static extern void FreeCString(IntPtr ptr);  

    static void Main() {  
        IntPtr result = RunValidationInput("some_input", "v1.0");  
        Console.WriteLine(Marshal.PtrToStringAnsi(result));  
        FreeCString(result);  
    }  
}  
```  

### **7. Swift**  

```swift
import Foundation  

let lib = dlopen("alembica.so", RTLD_LAZY)  
let function = dlsym(lib, "RunExtraction")  

if let f = function {  
    typealias ExtractFunction = @convention(c) (UnsafePointer<CChar>) -> UnsafePointer<CChar>  
    let extractFunc = unsafeBitCast(f, to: ExtractFunction.self)  
    let result = extractFunc("some_input")  
    print("Extraction Result:", String(cString: result))  
}  
```  

### **8. Node.js (ffi-napi)**  

```javascript
const ffi = require("ffi-napi");  

const lib = ffi.Library("./alembica.so", {  
    "RunComputeCosts": ["string", ["string"]]  
});  

const result = lib.RunComputeCosts("some_input");  
console.log("Cost Computation:", result);  
```  

### **9. Ruby (ffi gem)**  

```ruby
require 'ffi'  

module SharedLib  
  extend FFI::Library  
  ffi_lib './alembica.so'  

  attach_function :RunValidationInput, [:string, :string], :string  
end  

result = SharedLib.RunValidationInput("some_input", "v1.0")  
puts "Validation Result: #{result}"  
```  

### **10. PHP (FFI)**  

```php
$ffi = FFI::cdef("  
    char* RunComputeCosts(char* input);  
    void FreeCString(char* str);  
", "alembica.so");  

$result = $ffi->RunComputeCosts("some_input");  
echo "Cost Computation: $result\n";  
```  

### **11. Perl (FFI::Platypus)**  

```perl
use FFI::Platypus;  

my $ffi = FFI::Platypus->new();  
$ffi->lib("./alembica.so");  

$ffi->attach( 'RunExtraction' => ['string'] => 'string' );  

my $result = RunExtraction("some_input");  
print "Extraction Result: $result\n";  
```  

### **12. Kotlin (JNI)**  

```java
class AlembicaLib {  
    external fun RunComputeCosts(input: String): String  

    companion object {  
        init { System.loadLibrary("alembica") }  
    }  
}  

fun main() {  
    val lib = AlembicaLib()  
    val result = lib.RunComputeCosts("some_input")  
    println("Cost Computation: $result")  
}  
```  

### **13. Dart (dart:ffi)**  

```Dart
import 'dart:ffi';  
import 'dart:io';  

final dylib = DynamicLibrary.open("alembica.so");  

typedef RunComputeCosts_C = Pointer<Utf8> Function(Pointer<Utf8>);  
final RunComputeCosts =  
    dylib.lookupFunction<RunComputeCosts_C, RunComputeCosts_C>('RunComputeCosts');  

void main() {  
    final input = "some_input".toNativeUtf8();  
    final result = RunComputeCosts(input);  
    print("Cost Computation: ${result.toDartString()}");  
}  
```  

### **14. Fortran (ISO_C_BINDING)**  

```fortran
program test_alembica  
    use, intrinsic :: iso_c_binding  
    implicit none  

    interface  
        function RunComputeCosts(input) bind(C, name="RunComputeCosts")  
            import :: c_char  
            character(kind=c_char), dimension(*) :: RunComputeCosts  
        end function RunComputeCosts  
    end interface  

    character(kind=c_char), dimension(100) :: result  
    result = RunComputeCosts("some_input"//c_null_char)  

    print *, "Cost Computation:", result  
end program test_alembica  
```  

### **15. VBA (Declare Function - Windows Only)**  

```vb
Declare PtrSafe Function RunExtraction Lib "alembica.dll" (ByVal input As String) As String  

Sub TestAlembica()  
    Dim result As String  
    result = RunExtraction("some_input")  
    MsgBox "Extraction Result: " & result  
End Sub  
```  

### **16. Julia (ccall FFI)**  

```julia
function run_extraction(input::String)  
    result = ccall((:RunExtraction, "alembica.so"), Cstring, (Cstring,), input)  
    extracted = unsafe_string(result)  
    println("Extraction Result: ", extracted)  
end  

run_extraction("some_input")  
```  

### **17. R (Rcpp FFI)**  

```r
library(Rcpp)  

cppFunction('  
    #include <Rcpp.h>  
    extern "C" char* RunComputeCosts(const char* input);  

    // Wrapper function  
    std::string run_compute_costs(std::string input) {  
        return std::string(RunComputeCosts(input.c_str()));  
    }  
', includes = '#include "alembica.h"',  
   libs = "-L. -lalembica")  

# Call the function  
result <- run_compute_costs("some_input")  
print(paste("Cost Computation:", result))  
```  

## Conclusion  

The **`alembica.so`** shared library can be used in **many programming languages** that support **FFI (Foreign Function Interface)**.  
Beyond the languages covered here, it can also work with **other languages** such as **Tcl, Ada, and more**, provided they support calling C libraries.  

For more details, refer to the FFI documentation of your specific language.  

<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>