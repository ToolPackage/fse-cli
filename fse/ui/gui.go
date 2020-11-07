package ui

import (
	"github.com/ToolPackage/fse-cli/fse/client"
	"log"
)
import "github.com/jroimartin/gocui"

var gui *Gui

type Gui struct {
	ui *gocui.Gui
}

func Run() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	gui = &Gui{ui: g}

	c := client.NewClient()
	defer c.Close()

	g.Cursor = true
	g.Mouse = false
	g.SetManagerFunc(Layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err = g.MainLoop(); err != nil || err != gocui.ErrQuit {
		log.Fatal(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
