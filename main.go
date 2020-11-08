package main

import (
	"fmt"
	"github.com/ToolPackage/fse-cli/fse"
	"github.com/ToolPackage/fse-cli/fse/version"
	"github.com/eiannone/keyboard"
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	//version.Build = Build
	fmt.Print(version.SimpleLogo())
	fse.NewLine()
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if err = fse.InputHandler(char, key); err != nil {
			if err == fse.QuitError {
				break
			}
			panic(err)
		}
	}
}
