package main

import (
	"flag"

	"github.com/catplanet007/lit/lversion"
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "Show version information")
	flag.Parse()

	if showVersion {
		lversion.PrintVersionAndExit()
	}
}
