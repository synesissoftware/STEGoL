/*
 * Copyright (c) 2025 Matt Wilson and Synesis Information Systems
 *
 * Distributed under the 3-Clause BSD License (aka "New BSD-3 License"). See
 * accompanying file LICENSE file for details.
 */

package stegol_test

import (
	. "github.com/synesissoftware/STEGoL"

	"github.com/stretchr/testify/require"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 2
	Expected_VersionPatch uint16 = 3
	Expected_VersionAB    uint16 = 0x8001
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, VersionMajor)
	require.Equal(t, Expected_VersionMinor, VersionMinor)
	require.Equal(t, Expected_VersionPatch, VersionPatch)
	require.Equal(t, Expected_VersionAB, VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0002_0003_8001), Version)
}

func Test_Version_String(t *testing.T) {
	CheckStringEqual(t, "0.2.3-beta1", VersionString())
	require.Equal(t, "0.2.3-beta1", VersionString())
}
