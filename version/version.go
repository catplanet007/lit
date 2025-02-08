package version

import (
	"fmt"
	"os"
)

var (
	GitCommit string
	BuildTime string
	BuildTag  string
)

func PrintVersion(showVersion bool) {
	if showVersion {
		fmt.Printf("GitCommit: %s\n", GitCommit)
		fmt.Printf("BuildTime: %s\n", BuildTime)
		fmt.Printf("BuildTag: %s\n", BuildTag)
		os.Exit(0)
	}
}
