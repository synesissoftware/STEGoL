package stegol_test

import (
	stegol "github.com/synesissoftware/STEGoL"

	"testing"
)

type StrongBool bool

func Test_IsFalse(t *testing.T) {

	stegol.IsFalse(t, false)

	stegol.IsFalse(t, StrongBool(false))

	stegol.IsFalse(t, func() bool {
		return false
	})

	stegol.IsFalse(t, func() (bool, string) {
		return false, "oops"
	})
}

func Test_IsTrue(t *testing.T) {

	stegol.IsTrue(t, true)

	stegol.IsTrue(t, StrongBool(true))

	stegol.IsTrue(t, func() bool {
		return true
	})

	stegol.IsTrue(t, func() (bool, string) {
		return true, "oops"
	})
}
