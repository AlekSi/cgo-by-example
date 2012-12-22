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

func F() {
	fmt.Println("Go F()")
}

func main() {
	// constant
	C.f1(1)

	// cast Go int to C int
	var i int = 2
	C.f1(C.int(i))

	// convert Go string to C string (char*) and back
	cs := C.CString("Go string")
	csRet := C.f2(cs)
	fmt.Printf("fmt: %s\n", C.GoString(csRet))
	C.free(unsafe.Pointer(cs))
	C.free(unsafe.Pointer(csRet))

	// pass struct by value
	s1 := C.struct_s1{5} // or cast with C.int: s1 := C.struct_s1{C.int(i)}
	s1Ret := C.f31(s1)
	fmt.Printf("f31: s1=%v, s1Ret=%v\n", s1, s1Ret)

	// pass struct by pointer
	s1 = C.struct_s1{5}
	s1Ret = *C.f32(&s1)
	fmt.Printf("f32: s1=%v, s1Ret=%v\n", s1, s1Ret)

	// pass struct with int pointer
	// s2 := C.struct_s2{&C.int(i)} - cannot take the address of _Ctype_int(i), create variable as below
	ci := C.int(i)
	s2 := C.struct_s2{&ci}
	fmt.Printf("s2=%v, *s2.p=%d, i=%d\n", s2, *s2.p, i)
	C.f4(s2)
	fmt.Printf("f4: s2=%v, *s2.p=%d, i=%d\n", s2, *s2.p, i)

	f := func() {
		fmt.Println("Go f()")
	}

	fp := unsafe.Pointer(&f)
	// Pass Go function pointer to C and call it. Compiles but explodes. Do not do this!
	// C.f5((*[0]byte)(fp))

	f = F
	fp = unsafe.Pointer(&f)
	// Global function - same explosion. Do not do this!
	// C.f5((*[0]byte)(fp))

	_ = fp
}
