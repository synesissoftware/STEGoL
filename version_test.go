package stegol_test

import (
	"github.com/stretchr/testify/require"
	stegol "github.com/synesissoftware/STEGoL"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 2
	Expected_VersionPatch uint16 = 1
	Expected_VersionAB    uint16 = 0xFFFF
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, stegol.VersionMajor)
	require.Equal(t, Expected_VersionMinor, stegol.VersionMinor)
	require.Equal(t, Expected_VersionPatch, stegol.VersionPatch)
	require.Equal(t, Expected_VersionAB, stegol.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0002_0001_FFFF), stegol.Version)
}

func Test_Version_String(t *testing.T) {
	stegol.CheckStringEqual(t, "0.2.1", stegol.VersionString())
	require.Equal(t, "0.2.1", stegol.VersionString())
}
