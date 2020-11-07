package main

import (
	"github.com/ToolPackage/fse-cli/fse/ui"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//version.Build = Build

	ui.Run()
}
