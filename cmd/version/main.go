package main

import (
	"flag"

	"github.com/catplanet007/lit/version"
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "Show version information")
	flag.Parse()

	version.PrintVersion(showVersion)
}
