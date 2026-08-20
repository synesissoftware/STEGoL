package main

import (
	"fmt"

	stegol "github.com/synesissoftware/STEGoL"
	ver2go "github.com/synesissoftware/ver2go"
)

func main() {
	fmt.Printf("stegol v%s\n", stegol.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}
