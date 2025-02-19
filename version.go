package main

import (
	"runtime/debug"
)

// Version is this tool's semantic version number.
var Version string

func init() {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		Version = bi.Main.Version
	} else {
		Version = "undefined"
	}
}
