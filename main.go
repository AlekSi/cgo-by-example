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

	s1 := C.struct_s1{5}
	s1Ret := C.f31(s1)
	fmt.Printf("f31: s1=%#v, s1Ret=%#v\n", s1, s1Ret)

	s1 = C.struct_s1{5}
	s1Ret = *C.f32(&s1)
	fmt.Printf("f32: s1=%#v, s1Ret=%#v\n", s1, s1Ret)

	p := unsafe.Pointer(&i)
	s2 := C.struct_s2{p}
	s2Ret := C.f4(s2)
	fmt.Printf("f4: s2=%#v, s2Ret=%#v\n", s2, s2Ret)
}
