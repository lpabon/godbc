package godbc

import (
	//	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequire(t *testing.T) {
	a := 0
	Require(a != 0)
	/*
		assert.Panics(t, func() {
			Require(a == b)
		})

		assert.NotPanics(t, func() {
			Require(a != b)
		})

		assert.Panics(t, func() {
			Require(a > 0)
		})
	*/
}
