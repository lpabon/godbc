//+build prod

//
// Copyright 2014 Luis Pabon, Jr.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package godbc

// InvariantSimpleTester is an interface which provides a receiver to
// test the object
type InvariantSimpleTester interface {
	Invariant() bool
}

// InvariantTester is an interface which provides not only an Invariant(),
// but also a receiver to print the structure
type InvariantTester interface {
	InvariantSimpleTester
	String() string
}

// Require checks that the preconditions are satisfied before
// executing the function
//
// Example Code
//
// 		func Divide(a, b int) int {
//			godbc.Require(b != 0)
//			return a/b
// 		}
//
func Require(b bool, message ...interface{}) {
}

// Ensure checks the postconditions are satisfied before returning
// to the caller.
//
// Example Code
//
//		type Data struct {
//			a int
//		}
//
// 		func (*d Data) Set(a int) {
//			d.a = a
//			godbc.Ensure(d.a == a)
// 		}
//
func Ensure(b bool, message ...interface{}) {
}

// Check provides a simple assert
func Check(b bool, message ...interface{}) {
}

// InvariantSimple calls the objects Invariant() receiver to test
// the object for correctness.
//
// The caller object must provide an object that supports the
// interface InvariantSimpleTester and does not need to provide
// a String() receiver
func InvariantSimple(obj InvariantSimpleTester, message ...interface{}) {
}

// Invariant calls the objects Invariant() receiver to test
// the object for correctness.
//
// The caller object must provide an object that supports the
// interface InvariantTester
//
// To see an example, please take a look at the godbc_test.go
func Invariant(obj InvariantTester, message ...interface{}) {
}
