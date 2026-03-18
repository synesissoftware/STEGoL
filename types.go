// Copyright 2018-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 17th March 2026
 * Updated: 18th March 2026
 */

package stegol

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type SignedIntegerConvertible interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedIntegerConvertible interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type IntegerConvertible interface {
	SignedIntegerConvertible | UnsignedIntegerConvertible
}

type BooleanConvertible interface {
	~bool
}

type Predicate = func () bool

type PredicateWithFailureMessage = func () (bool, string)
