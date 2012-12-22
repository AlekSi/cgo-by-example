package main

/*
#include "lib/lib.h"
#include <stdlib.h>       // for free()
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
	// Simple C function call.
	C.f0()

	// Call with constant.
	C.f1(1)

	// Call with Go int - cast to C int.
	var i int = 2
	C.f1(C.int(i))

	// Convert Go string to C string (char*) and back.
	cs := C.CString("Go string")
	csRet := C.f2(cs)
	fmt.Printf("fmt: %s\n", C.GoString(csRet))
	C.free(unsafe.Pointer(cs))          // free memory right now ...
	defer C.free(unsafe.Pointer(csRet)) // ... or later

	// Pass C struct to C function by value.
	s1 := C.struct_s1{5} // or cast with C.int: s1 := C.struct_s1{C.int(i)}
	s1Ret := C.f31(s1)
	fmt.Printf("f31: s1=%v, s1Ret=%v\n", s1, s1Ret)

	// Pass C struct to C function by pointer.
	s1 = C.struct_s1{5}
	s1Ret = *C.f32(&s1)
	fmt.Printf("f32: s1=%v, s1Ret=%v\n", s1, s1Ret)

	// Pass C struct with int pointer to C function.
	// s2 := C.struct_s2{&C.int(i)} - compile error "cannot take the address of _Ctype_int(i)", create variable as below
	ci := C.int(i)
	s2 := C.struct_s2{&ci}
	fmt.Printf("s2=%v, *s2.p=%d, i=%d\n", s2, *s2.p, i)
	C.f4(s2)
	fmt.Printf("f4: s2=%v, *s2.p=%d, i=%d\n", s2, *s2.p, i)

	// Pass C function pointer to C function.
	// fp := unsafe.Pointer(C.f0) - compile error "must call C.f0"

	f := func() {
		fmt.Println("Go f()")
	}
	fp := unsafe.Pointer(&f)
	// Pass Go function pointer to C function. Compiles but explodes. Do not do this!
	// C.f5((*[0]byte)(fp))

	f = F
	fp = unsafe.Pointer(&f)
	// Pass global Go function pointer to C function. Compiles but explodes. Do not do this!
	// C.f5((*[0]byte)(fp))

	_ = fp
}
