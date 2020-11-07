package main

import (
	"github.com/ToolPackage/fse-cli/fse"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//version.Build = Build

	fse.Run()
}
