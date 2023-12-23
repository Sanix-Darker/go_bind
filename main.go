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
