package main

/*
#cgo pkg-config: python3
#define Py_LIMITED_API
#include <Python.h>

*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/golang/protobuf/proto"
)

//export GetDataInfo
func GetDataInfo(pb *C.char) *C.char {
	// Convert C string to Go byte slice
	bytes := C.GoBytes(
		unsafe.Pointer(pb),
		C.int(C.strlen(pb)),
	)
	data := &Pack{}

	// Unmarshal the protobuf data
	err := proto.Unmarshal(bytes, data)
	if err != nil {
		fmt.Printf("Error unmarshaling protobuf: %v", err)
		return nil
	}

	// Print information about the data
	fmt.Printf("Data: %s, Length: %d\n", data.GetData(), data.GetLen())

	// Manipulate the data
	data.Data = "Hi from GO, your input was '" + data.GetData() + "'"
	data.Len = int32(len(data.GetData()))

	// Marshal the modified protobuf data
	modifiedBytes, err := proto.Marshal(data)
	if err != nil {
		fmt.Printf("Error marshaling protobuf: %v", err)
		return nil
	}

	// Convert Go byte slice to C string
	return C.CString(string(modifiedBytes))
}

func main() {}