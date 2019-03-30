/* /////////////////////////////////////////////////////////////////////////
 * File:        check_string.go
 *
 * Purpose:     CheckString*() functions
 *
 * Created:     2nd April 2018
 * Updated:     30th March 2019
 *
 * Copyright (c) 2018-2019, Matthew Wilson and Synesis Software
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * - Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 * - Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in the
 *   documentation and/or other materials provided with the distribution.
 * - Neither the names of Matthew Wilson, Sean Kelly, Synesis Software nor
 *   the names of any contributors may be used to endorse or promote products
 *   derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * ////////////////////////////////////////////////////////////////////// */

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

func extract_flags(options ...interface{}) (result CheckStringFlag) {

	for _, arg := range options {

		switch v := arg.(type) {

		case CheckStringFlag:

			result |= v
		default:

			// Default is to ignore the rest
			;
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
			;
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

// CheckStringEqual() evaluates two strings, calling testing.T.Errorf() if
// the evaluation fails. The evaluation is done by equality comparison.
func CheckStringEqual(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return e.(string) == a }, "actual value", "is not equal to expected value", "", options...)
}

// CheckStringEqualTrimmed() evaluates two strings, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equality comparison after whitespace-trimming the actual value.
func CheckStringEqualTrimmed(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return strings.TrimSpace(e.(string)) == strings.TrimSpace(a) }, "actual value", "when trimmed, is different to expected value", "", options...)
}

// CheckStringEqualIgnoreCase() evaluates two strings, calling
// testing.T.Errorf() if the evaluation fails. The evaluation is done by
// equalitycomparison via the strings.EqualFold() standard library function.
func CheckStringEqualIgnoreCase(t *testing.T, expected, actual string, options ...interface{}) {

	t.Helper()

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return strings.EqualFold(e.(string), a) }, "actual value", "is different, when ignoring case, to expected value", "", options...)
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

	checkStringCompare(t, expected, actual, func(e interface{}, a string) bool { return fn(e.(string), a) }, "actual value", "does not compare equal to expected value", "when compared by " + comparison_type, options...)
}

// CheckStringEqualAny() evaluates a string value against an array of
// expected string values, calling testing.T.Errorf() if the evaluation
// fails. The evaluation is done by equality comparison
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


