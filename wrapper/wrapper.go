package wrapper

/*
#include "wrapper.h"
#cgo LDFLAGS: -L. -llib
*/
import "C"

import (
	"fmt"
)

//export F
func F(i C.int) {
	fmt.Printf("Go F(%d)\n", i)
}

func CallF5WithF() {
	C.call_f5_with_F()
}
