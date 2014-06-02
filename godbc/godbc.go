package godbc

import (
	"fmt"
	"runtime"
)

func Require(b bool, message ...interface{}) {
	if !b {
		pc, file, line, ok := runtime.Caller(1)
		finfo := runtime.FuncForPC(pc)
		fmt.Println("Hello")
		fmt.Printf("Func %s - %s - %d - %t",
			finfo.Name(),
			file,
			line,
			ok)
		panic("Failed---")
	}
}

func Ensure(b bool) {
	if !b {
		panic("Ensure context failed")
	}
}

func Check() {
	fmt.Println("In Check")
}

func Invariant() {
	fmt.Println("In Inv")
}
