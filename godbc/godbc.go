package godbc

import (
	"errors"
	"fmt"
	"runtime"
)

func dbc_panic(dbc_func_name string, b bool, message ...interface{}) {
	if !b {

		// Get caller information which is the caller
		// of the caller of this function
		pc, file, line, _ := runtime.Caller(2)
		caller_func_info := runtime.FuncForPC(pc)

		error_string := fmt.Sprintf("%s: %s:%s:%d",
			dbc_func_name,
			file,
			caller_func_info.Name(),
			line)
		if len(message) > 0 {
			error_string += fmt.Sprintf("\n\rMessage:%+v len %d", message, len(message))
		}
		err := errors.New(error_string)

		// Finally panic
		panic(err)
	}
}

func Require(b bool, message ...interface{}) {
	dbc_panic("REQUIRE", b, message...)
}

func Ensure(b bool, message ...interface{}) {
	dbc_panic("ENSURE", b, message...)
}

func Check(b bool, message ...interface{}) {
	dbc_panic("CHECK", b, message...)
}
