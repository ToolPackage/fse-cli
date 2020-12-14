package main

import (
	"github.com/ToolPackage/fse-cli/internal"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	_, err := flags.Parse(&internal.Opts)
	if err != nil {
		os.Exit(1)
	}
}
