// Copyright 2018-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stegol

import (
	"testing"
	"unsafe"
)

func _obtainBoolValueOrPanic[T any](value *T) bool {

	if unsafe.Sizeof(*value) != unsafe.Sizeof(bool(false)) {
		panic("")
	}

	return *(*bool)(unsafe.Pointer(value))
}

// Evaluates a boolean value or predicate to obtain a boolean value, calling
// testing.T.Errorf() if the value is not false.
func IsFalse[T bool | BooleanConvertible | Predicate | PredicateWithFailureMessage](t *testing.T, value T) {

	t.Helper()

	var passes bool
	var failureMessage string

	switch v := any(value).(type) {
	case Predicate:

		var b bool
		b = v()
		passes = !b
	case PredicateWithFailureMessage:

		var b bool
		b, failureMessage = v()
		passes = !b
	default:

		// this is valid ONLY because we permit ONLY `BooleanConvertible` and
		// `bool` in addition to the predicates

		b := _obtainBoolValueOrPanic(&value)
		passes = !b
	}

	if !passes {
		if failureMessage != "" {

			t.Errorf(failureMessage)
		} else {

			t.Errorf("expected value to be false")
		}
	}
}

// Evaluates a boolean value or predicate to obtain a boolean value, calling
// testing.T.Errorf() if the value is not true.
func IsTrue[T bool | BooleanConvertible | Predicate | PredicateWithFailureMessage](t *testing.T, value T) {

	t.Helper()

	var passes bool
	var failureMessage string

	switch v := any(value).(type) {
	case Predicate:

		var b bool
		b = v()
		passes = b
	case PredicateWithFailureMessage:

		var b bool
		b, failureMessage = v()
		passes = b
	default:

		// this is valid ONLY because we permit ONLY `BooleanConvertible` and
		// `bool` in addition to the predicates

		b := _obtainBoolValueOrPanic(&value)
		passes = b
	}

	if !passes {
		if failureMessage != "" {

			t.Errorf(failureMessage)
		} else {

			t.Errorf("expected value to be true")
		}
	}
}
