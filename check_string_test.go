
package stegol_test

import (

	stegol "github.com/synesissoftware/STEGoL"

	"testing"
)

func Test_CheckStringEqual(t *testing.T) {

	stegol.CheckStringEqual(t, "", "")
	stegol.CheckStringEqual(t, "a", "a")
	stegol.CheckStringEqual(t, "\tabc", "\tabc")
}

func Test_CheckStringNotEqual(t *testing.T) {

	stegol.CheckStringNotEqual(t, "", " ")
	stegol.CheckStringNotEqual(t, "a", "a\t")
	stegol.CheckStringNotEqual(t, "abc", "\tabc")
}

func Test_CheckStringEqualChomped(t *testing.T) {

	stegol.CheckStringEqualChomped(t, "", "")
	stegol.CheckStringEqualChomped(t, "a", "a")
	stegol.CheckStringEqualChomped(t, "\tabc", "\tabc")

	stegol.CheckStringEqualChomped(t, "", "\n")
	stegol.CheckStringEqualChomped(t, "\n", "\n\n")
	stegol.CheckStringEqualChomped(t, "", "\r")
	stegol.CheckStringEqualChomped(t, "\n", "\n\r")
	stegol.CheckStringEqualChomped(t, "", "\r\n")

	stegol.CheckStringEqualChomped(t, "abc", "abc\n")
	stegol.CheckStringEqualChomped(t, "abc\n", "abc\n\n")
	stegol.CheckStringEqualChomped(t, "abc", "abc\r")
	stegol.CheckStringEqualChomped(t, "abc\n", "abc\n\r")
	stegol.CheckStringEqualChomped(t, "abc", "abc\r\n")
}

func Test_CheckStringNotEqualChomped(t *testing.T) {

	stegol.CheckStringNotEqualChomped(t, "a", "ab")
	stegol.CheckStringNotEqualChomped(t, "\tabc", "\tabcd")

	stegol.CheckStringNotEqualChomped(t, "", "\na")
	stegol.CheckStringNotEqualChomped(t, "\n", "\n\na")
	stegol.CheckStringNotEqualChomped(t, "", "\ra")
	stegol.CheckStringNotEqualChomped(t, "\n", "\n\ra")
	stegol.CheckStringNotEqualChomped(t, "", "\r\na")

	stegol.CheckStringNotEqualChomped(t, "abc", "abc\n\n")
	stegol.CheckStringNotEqualChomped(t, "abc", "abc\r\r")
	stegol.CheckStringNotEqualChomped(t, "abc", "abc\n\r")
}

func Test_CheckStringEqualTrimmed(t *testing.T) {

	stegol.CheckStringEqual(t, "", "")
	stegol.CheckStringEqual(t, "a", "a")
	stegol.CheckStringEqual(t, "\tabc", "\tabc")

	stegol.CheckStringEqualTrimmed(t, "", " ")
	stegol.CheckStringEqualTrimmed(t, "a", "a\t")
	stegol.CheckStringEqualTrimmed(t, "abc", "\tabc")
}

