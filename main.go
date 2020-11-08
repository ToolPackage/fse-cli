package main

import (
	"github.com/ToolPackage/fse-cli/fse"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//version.Build = Build

	fse.Client = fse.NewClient()
	defer fse.Client.Close()

	fse.UI = new(fse.Gui)

	fse.RunGui()
}
