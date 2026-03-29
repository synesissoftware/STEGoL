/*
 * Copyright (c) 2025-2026 Matt Wilson and Synesis Information Systems
 *
 * Distributed under the 3-Clause BSD License (aka "New BSD-3 License"). See
 * accompanying file LICENSE file for details.
 */

package stegol_test

import (
	. "github.com/synesissoftware/STEGoL"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 3
	Expected_VersionPatch uint16 = 0
	Expected_VersionAB    uint16 = 0xFFFF
)

func Test_Version_Elements(t *testing.T) {
	CheckIntegerEqual(t, Expected_VersionMajor, VersionMajor)
	CheckIntegerEqual(t, Expected_VersionMinor, VersionMinor)
	CheckIntegerEqual(t, Expected_VersionPatch, VersionPatch)
	CheckIntegerEqual(t, Expected_VersionAB, VersionAB)
}

func Test_Version(t *testing.T) {
	CheckIntegerEqual(t, 0x0000_0003_0000_FFFF, Version)
}

func Test_Version_String(t *testing.T) {
	CheckStringEqual(t, "0.3.0", VersionString())
}
