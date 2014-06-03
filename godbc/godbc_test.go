package godbc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequire(t *testing.T) {
	a, b := 0, 1
	assert.Panics(t, func() {
		Require(a == b)
	})

	assert.NotPanics(t, func() {
		Require(a != b)
	})

	assert.Panics(t, func() {
		Require(a > 0)
	})
}

func TestEnsure(t *testing.T) {
	a, b := 0, 1
	assert.Panics(t, func() {
		Ensure(a == b)
	})

	assert.NotPanics(t, func() {
		Ensure(a != b)
	})

	assert.Panics(t, func() {
		Ensure(a > 0)
	})
}

func TestCheck(t *testing.T) {
	a, b := 0, 1
	assert.Panics(t, func() {
		Check(a == b)
	})

	assert.NotPanics(t, func() {
		Check(a != b)
	})

	assert.Panics(t, func() {
		Check(a > 0)
	})
}
