package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/danielmiessler/fabric/internal/cli"
	"github.com/danielmiessler/fabric/internal/pipeline"
)

func main() {
	if handled, err := pipeline.CleanupRunDirFromEnv(); handled {
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		return
	}

	err := cli.Cli(version)
	if err != nil && !flags.WroteHelp(err) {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
