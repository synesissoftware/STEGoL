// Copyright 2018-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 2nd April 2018
 * Updated: 23rd February 2025
 */

package stegol

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

type CheckStringFlag int

const (
	CheckStringFlag_None = 1 << iota
)

type StringCompareFunc func(expected, actual string) bool

type stringCompareFunc func(expected interface{}, actual string) bool

/* /////////////////////////////////////////////////////////////////////////
 * helper functions
 */

func chomp_string(s string) string {

	switch l := len(s); l {

	case 0:

		break
	default:

		if '\r' == s[l-2] && '\n' == s[l-1] {

			s = s[0 : l-2]
			break
		}

		fallthrough
	case 1:

		switch s[l-1] {

		case '\r', '\n':

			s = s[0 : l-1]
		}
	}

	return s
}

func extract_flags(options ...interface{}) (result CheckStringFlag) {

	for _, arg := range options {

		switch v := arg.(type) {

		case CheckStringFlag:

			result |= v
		default:

			// Default is to ignore the rest
		}
	}

	return
}

func checkStringCompare(t *testing.T, expected interface{}, actual string, fn stringCompareFunc, frag0, frag1, frag2 string, options ...interface{}) {

	flags := extract_flags(options...)

	_ = flags

	t.Helper()

	if !fn(expected, actual) {

		do_multiline := false

		if !do_multiline {

			switch exp_s := expected.(type) {

			case string:

				if strings.Index(exp_s, "\n") >= 0 || strings.Index(actual, "\n") >= 0 {

					do_multiline = true
				}
			default:

				do_multiline = true
			}

		}

		format := ""

		if do_multiline {

			format = fmt.Sprintf("%s\n\t'%%v'\n%s\n\t'%%v'\n%s", frag0, frag1, frag2)
		} else {

			format = fmt.Sprintf("%s '%%v' %s '%%v'%s", frag0, frag1, frag2)
		}

		t.Errorf(format, actual, expected)
	}
}

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */

// CheckStringEqual() evaluates two strings for equality, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equality comparison.
func CheckStringEqual(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return e.(string) == a }, "actual value", "is not equal to expected value", "", options...)
}

// CheckStringNotEqual() evaluates two strings for inequality, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equality comparison.
func CheckStringNotEqual(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return e.(string) != a }, "actual value", "is not different to expected value", "", options...)
}

// CheckStringEqualChomped() evaluates two strings for equality, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equality comparison after chomping the actual value.
func CheckStringEqualChomped(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	actual = chomp_string(actual)

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return e.(string) == a }, "actual value (when chomped)", "is not equal to expected value", "", options...)
}

// CheckStringNotEqualChomped() evaluates two strings for inequality, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equality comparison after chomping the actual value.
func CheckStringNotEqualChomped(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	actual = chomp_string(actual)

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return e.(string) != a }, "actual value (when chomped)", "is not different to expected value", "", options...)
}

// CheckStringEqualTrimmed() evaluates two strings for equality, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equality comparison after whitespace-trimming the actual value.
func CheckStringEqualTrimmed(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return strings.TrimSpace(e.(string)) == strings.TrimSpace(a) }, "actual value", "when trimmed, is different to expected value", "", options...)
}

// CheckStringNotEqualTrimmed() evaluates two strings for inequality, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equality comparison after whitespace-trimming the actual value.
func CheckStringNotEqualTrimmed(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return strings.TrimSpace(e.(string)) == strings.TrimSpace(a) }, "actual value", "when trimmed, is not different to expected value", "", options...)
}

// CheckStringEqualIgnoreCase() evaluates two strings for equality, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equality comparison ignoring the case of the strings, via the
// strings.EqualFold() standard library function.
func CheckStringEqualIgnoreCase(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return strings.EqualFold(e.(string), a) }, "actual value", "is different, when ignoring case, to expected value", "", options...)
}

// CheckStringEqualIgnoreCase() evaluates two strings for inequality, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equality comparison ignoring the case of the strings, via the
// strings.EqualFold() standard library function.
func CheckStringNotEqualIgnoreCase(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return !strings.EqualFold(e.(string), a) }, "actual value", "is not different, when ignoring case, to expected value", "", options...)
}

func CheckStringByStringMatch(t *testing.T, pattern string, actual string, options ...interface{}) {

	t.Helper()

	if re, err := regexp.Compile(pattern); err != nil {

		t.Errorf("The given pattern - %q - could not be compiled as a regular expression: %v\n", pattern, err)
	} else {

		checkStringCompare(t, pattern, actual, func(e interface{}, a string) bool { return re.MatchString(a) }, "actual value", "does not match the regular expression", "", options...)
	}
}

// CheckStringCompare() evaluates two strings, calling testing.T.Errorf()
// if the evaluation fails. The evaluation is done by the given
// caller-supplied comparison function, whose brief descriptor
// comparison_type, e.g. "regular expression" will be prefixed with the
// string "when compared by ".
func CheckStringCompare(t *testing.T, expected, actual string, fn StringCompareFunc, comparison_type string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return fn(e.(string), a) }, "actual value", "does not compare equal to expected value", "when compared by "+comparison_type, options...)
}

// CheckStringEqualAny() evaluates a string for equality against an array of
// string values, calling testing.T.Errorf() if every evaluation fails. Each
// evaluation is done by equality comparison.
func CheckStringEqualAny(t *testing.T, expecteds []string, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expecteds, actual, func(e interface{}, a string) bool {

		for _, expected := range expecteds {

			if expected == a {

				return true
			}
		}

		return false
	}, "actual value", "is not equal to any of the expected values", "", options...)
}

/* ///////////////////////////// end of file //////////////////////////// */
