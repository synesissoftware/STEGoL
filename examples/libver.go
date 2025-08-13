package main

import (
	stegol "github.com/synesissoftware/STEGoL"
	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("stegol v%s\n", ver2go.CalcVersionString(stegol.VersionMajor, stegol.VersionMinor, stegol.VersionPatch, stegol.VersionAB))
	fmt.Printf("ver2go v%s\n", ver2go.CalcVersionString(ver2go.VersionMajor, ver2go.VersionMinor, ver2go.VersionPatch, ver2go.VersionAB))
}
