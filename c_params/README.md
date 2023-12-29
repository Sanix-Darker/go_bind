## GO-BIND

How much do you think it's cool to code in golang and use bindings C-lib from
an .so of your functions in a python code ?
In this repo, am just illustrating an example on how to do that maggic.

From a go code :
```go
package main

/*
#cgo pkg-config: python3
#define Py_LIMITED_API
#include <Python.h>

typedef struct {
    int xxx;
    char yyy[100];
} GoBindStruct;

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type GoBindStruct struct {
	xxx C.int
	yyy *C.char
}

//export PrintThis
func PrintThis(a C.int, b *C.char) C.int {
	aa := GoBindStruct{
		a, b,
	}
	fmt.Printf("> input :: %d\n", aa.xxx)
	fmt.Printf("< output :: %s\n", C.GoString(aa.yyy))
	// aa.xxx = aa.xxx + C.int(unsafe.Sizeof(aa.yyy))
	aa.xxx = aa.xxx + C.int(len(C.GoString(aa.yyy)))

	return aa.xxx
}

func main() {
	var strr *C.char = C.CString("dk")
	defer C.free(unsafe.Pointer(strr))

	PrintThis(1, strr)
}
```
To a python code :

```python
from go_bind_py import lib

print("=", lib.PrintThis(4, b"tintin"))
print(f">> go_bind_py :: export {dir(lib)}")

#
# > input :: 4
# < output :: tintin
# = 10
# >> go_bind_py :: export ['PrintThis']
```

## WHY

So... you may wondering.. WHY DO I CARE !?
actually you should care; python functions & objects are HEAVY when they are
wrote well... in python code, if you want to optimize your python code, the
workaround is to deal with a C-shared lib, from any other language, in this
case, golang.

**NOTE:** This is only if you're doing complex or __HARD__ CPU/RAM
processing... normally a CRUD doesn't need this sorcery.

## HOW

Am using here FFI, from cffi lib,

FFI: `Foreign function interface`

A foreign function interface (FFI) is a mechanism by which a program written in one programming
language can call routines or make use of services written or compiled in another one.
An FFI is often used in contexts where calls are made into binary dynamic-link library.

## REQUIREMENTS

- make
- golang (>1.18 recommended)
- python (>=3.11 recommended)
- cffi (from pypi)

## HOW TO BUILD

- First build the .so (i made the make file extremly generic btw): `make build-so`
- Then build the (.c, .so, .h) files that will help use your module functions
  inside your python code. `make build-py`
- You can test now your python code easily `make run` (note that this will run
 the .so builder and the .py ffi linker too).
