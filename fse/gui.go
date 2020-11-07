package fse

import (
	"github.com/ToolPackage/fse-cli/fse/ui"
	"log"
)
import "github.com/jroimartin/gocui"

func Run() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	client := NewClient(g)
	defer client.Close()

	g.Cursor = true
	g.Mouse = false
	g.SetManagerFunc(ui.Layout)

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
