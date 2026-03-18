// Copyright 2018-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 17th March 2026
 * Updated: 18th March 2026
 */

package stegol

import (
	"fmt"
	"testing"
)

func compare_int64_int64(
	lhs int64,
	rhs int64,
) int {

	if lhs < rhs {
		return -1
	}

	if lhs > rhs {
		return +1
	}

	return 0
}

func compare_uint64_uint64(
	lhs uint64,
	rhs uint64,
) int {

	if lhs < rhs {
		return -1
	}

	if lhs > rhs {
		return +1
	}

	return 0
}

func compare_int64_uint64(
	lhs int64,
	rhs uint64,
) int {

	if lhs < 0 {
		return -1
	}

	lhs_u := uint64(lhs)

	if lhs_u < rhs {
		return -1
	}

	if lhs_u > rhs {
		return +1
	}

	return 0
}

func compareInteger[T1 int64 | uint64, T2 int64 | uint64](
	lhs T1,
	rhs T2,
) int {
	switch any(lhs).(type) {
	case int64:

		switch any(rhs).(type) {
		case int64:

			return compare_int64_int64(int64(lhs), int64(rhs))
		case uint64:

			return compare_int64_uint64(int64(lhs), uint64(rhs))
		default:

			panic("VIOLATION: unexpected type")
		}
	case uint64:

		switch any(rhs).(type) {
		case int64:

			return -compare_int64_uint64(int64(rhs), uint64(lhs))
		case uint64:

			return compare_uint64_uint64(uint64(lhs), uint64(rhs))
		default:

			panic("VIOLATION: unexpected type")
		}
	default:

		panic("VIOLATION: unexpected type")
	}
}

func checkIntegerEqual[T1 Integer, T2 Integer](
	expected T1,
	actual T2,
) (
	passes bool,
	failureMessage string,
) {
	switch any(expected).(type) {
	case int, int8, int16, int32, int64:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(int64(expected), int64(actual)) == 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(int64(expected), uint64(actual)) == 0
		default:

			panic("VIOLATION: unexpected type")
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(uint64(expected), int64(actual)) == 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(uint64(expected), uint64(actual)) == 0
		default:

			panic("VIOLATION: unexpected type")
		}
	default:

		panic("VIOLATION: unexpected type")
	}

	if !passes {

		failureMessage = fmt.Sprintf("actual value %v not equal to the expected %v", actual, expected)
	}

	return
}

func checkIntegerGreater[T1 Integer, T2 Integer](
	expected T1,
	actual T2,
) (
	passes bool,
	failureMessage string,
) {
	switch any(expected).(type) {
	case int, int8, int16, int32, int64:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(int64(expected), int64(actual)) < 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(int64(expected), uint64(actual)) < 0
		default:

			panic("VIOLATION: unexpected type")
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(uint64(expected), int64(actual)) < 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(uint64(expected), uint64(actual)) < 0
		default:

			panic("VIOLATION: unexpected type")
		}
	default:

		panic("VIOLATION: unexpected type")
	}

	if !passes {

		failureMessage = fmt.Sprintf("actual value %v not greater than the expected %v", actual, expected)
	}

	return
}

func checkIntegerGreaterOrEqual[T1 Integer, T2 Integer](
	expected T1,
	actual T2,
) (
	passes bool,
	failureMessage string,
) {
	switch any(expected).(type) {
	case int, int8, int16, int32, int64:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(int64(expected), int64(actual)) <= 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(int64(expected), uint64(actual)) <= 0
		default:

			panic("VIOLATION: unexpected type")
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(uint64(expected), int64(actual)) <= 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(uint64(expected), uint64(actual)) <= 0
		default:

			panic("VIOLATION: unexpected type")
		}
	default:

		panic("VIOLATION: unexpected type")
	}

	if !passes {

		failureMessage = fmt.Sprintf("actual value %v not greater than or equal to the expected %v", actual, expected)
	}

	return
}

func checkIntegerLess[T1 Integer, T2 Integer](
	expected T1,
	actual T2,
) (
	passes bool,
	failureMessage string,
) {
	switch any(expected).(type) {
	case int, int8, int16, int32, int64:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(int64(expected), int64(actual)) > 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(int64(expected), uint64(actual)) > 0
		default:

			panic("VIOLATION: unexpected type")
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(uint64(expected), int64(actual)) > 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(uint64(expected), uint64(actual)) > 0
		default:

			panic("VIOLATION: unexpected type")
		}
	default:

		panic("VIOLATION: unexpected type")
	}

	if !passes {

		failureMessage = fmt.Sprintf("actual value %v not less than the expected %v", actual, expected)
	}

	return
}

func checkIntegerLessOrEqual[T1 Integer, T2 Integer](
	expected T1,
	actual T2,
) (
	passes bool,
	failureMessage string,
) {
	switch any(expected).(type) {
	case int, int8, int16, int32, int64:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(int64(expected), int64(actual)) >= 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(int64(expected), uint64(actual)) >= 0
		default:

			panic("VIOLATION: unexpected type")
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(uint64(expected), int64(actual)) >= 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(uint64(expected), uint64(actual)) >= 0
		default:

			panic("VIOLATION: unexpected type")
		}
	default:

		panic("VIOLATION: unexpected type")
	}

	if !passes {

		failureMessage = fmt.Sprintf("actual value %v not less than or equal to the expected %v", actual, expected)
	}

	return
}

func checkIntegerNotEqual[T1 Integer, T2 Integer](
	expected T1,
	actual T2,
) (
	passes bool,
	failureMessage string,
) {
	switch any(expected).(type) {
	case int, int8, int16, int32, int64:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(int64(expected), int64(actual)) != 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(int64(expected), uint64(actual)) != 0
		default:

			panic("VIOLATION: unexpected type")
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:

		switch any(actual).(type) {
		case int, int8, int16, int32, int64:

			passes = compareInteger(uint64(expected), int64(actual)) != 0
		case uint, uint8, uint16, uint32, uint64, uintptr:

			passes = compareInteger(uint64(expected), uint64(actual)) != 0
		default:

			panic("VIOLATION: unexpected type")
		}
	default:

		panic("VIOLATION: unexpected type")
	}

	if !passes {

		failureMessage = fmt.Sprintf("actual value %v equal to the expected %v", actual, expected)
	}

	return
}

// Evaluates two integers, of arbitrary type, for equality comparison,
// calling testing.T.Errorf() if the evaluation fails.
func CheckIntegerEqual[T1 Integer, T2 Integer](
	t *testing.T,
	expected T1,
	actual T2,
) {

	t.Helper()

	if passes, failureMessage := checkIntegerEqual(expected, actual); !passes {

		t.Error(failureMessage)
	}
}

// Evaluates two integers, of arbitrary type, for less-than comparison,
// calling testing.T.Errorf() if the evaluation fails.
func CheckIntegerLess[T1 Integer, T2 Integer](
	t *testing.T,
	expected T1,
	actual T2,
) {

	t.Helper()

	if passes, failureMessage := checkIntegerLess(expected, actual); !passes {

		t.Error(failureMessage)
	}
}

// Evaluates two integers, of arbitrary type, for less-than-or-equal
// comparison, calling testing.T.Errorf() if the evaluation fails.
func CheckIntegerLessOrEqual[T1 Integer, T2 Integer](
	t *testing.T,
	expected T1,
	actual T2,
) {

	t.Helper()

	if passes, failureMessage := checkIntegerLessOrEqual(expected, actual); !passes {

		t.Error(failureMessage)
	}
}

// Evaluates two integers, of arbitrary type, for greater-than comparison,
// calling testing.T.Errorf() if the evaluation fails.
func CheckIntegerGreater[T1 Integer, T2 Integer](
	t *testing.T,
	expected T1,
	actual T2,
) {

	t.Helper()

	if passes, failureMessage := checkIntegerGreater(expected, actual); !passes {

		t.Error(failureMessage)
	}
}

// Evaluates two integers, of arbitrary type, for greater-than-or-equal
// comparison, calling testing.T.Errorf() if the evaluation fails.
func CheckIntegerGreaterOrEqual[T1 Integer, T2 Integer](
	t *testing.T,
	expected T1,
	actual T2,
) {

	t.Helper()

	if passes, failureMessage := checkIntegerGreaterOrEqual(expected, actual); !passes {

		t.Error(failureMessage)
	}
}

// Evaluates two integers, of arbitrary type, for inequality comparison,
// calling testing.T.Errorf() if the evaluation fails.
func CheckIntegerNotEqual[T1 Integer, T2 Integer](
	t *testing.T,
	expected T1,
	actual T2,
) {

	t.Helper()

	if passes, failureMessage := checkIntegerNotEqual(expected, actual); !passes {

		t.Error(failureMessage)
	}
}
