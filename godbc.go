/*
 * Copyright 2014 Luis Pabon, Jr.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package godbc

import (
	"errors"
	"fmt"
	"runtime"
)

type InvariantSimpleTester interface {
	Invariant() bool
}

type InvariantTester interface {
	InvariantSimpleTester
	String() string
}

func dbc_panic(dbc_func_name string, b bool, message ...interface{}) {
	if !b {

		// Get caller information which is the caller
		// of the caller of this function
		pc, file, line, _ := runtime.Caller(2)
		caller_func_info := runtime.FuncForPC(pc)

		error_string := fmt.Sprintf("%s:\n\r\tfunc (%s) 0x%x\n\r\tFile %s:%d",
			dbc_func_name,
			caller_func_info.Name(),
			pc,
			file,
			line)

		if len(message) > 0 {
			error_string += fmt.Sprintf("\n\r\tInfo: %+v", message)
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

func InvariantSimple(obj InvariantSimpleTester, message ...interface{}) {
	dbc_panic("INVARIANT", obj.Invariant(), message...)
}

func Invariant(obj InvariantTester, message ...interface{}) {
	m := append(message, obj)
	dbc_panic("INVARIANT", obj.Invariant(), m)
}
