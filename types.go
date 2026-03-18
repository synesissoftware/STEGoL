// Copyright 2018-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 17th March 2026
 * Updated: 18th March 2026
 */

package stegol

type BooleanConvertible interface {
	~bool
}

type Predicate = func () bool

type PredicateWithFailureMessage = func () (bool, string)
