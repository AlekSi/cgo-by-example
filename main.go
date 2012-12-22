package main

/*
#include "lib/lib.h"
#cgo LDFLAGS: -L. -llib
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	C.f1(1)

	i := 2
	C.f1(C.int(i))

	cs := C.CString("Go string")
	csRet := C.f2(cs)
	fmt.Printf("fmt: %s\n", C.GoString(csRet))
	C.free(unsafe.Pointer(cs))
	C.free(unsafe.Pointer(csRet))
}
