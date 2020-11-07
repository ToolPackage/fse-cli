package main

import (
	"github.com/ToolPackage/fse-cli/fse"
	"github.com/ToolPackage/fse-cli/fse/ui"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//version.Build = Build

	fse.Client = fse.NewClient()
	defer fse.Client.Close()

	ui.UI = new(ui.Gui)

	ui.RunGui()
}
