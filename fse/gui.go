package fse

import (
	"github.com/0xAX/notificator"
	"github.com/ToolPackage/fse-cli/fse/version"
	"github.com/jroimartin/gocui"
	"log"
)

func RunGui() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	Notify = notificator.New(notificator.Options{
		AppName: version.Name,
	})
	UI.SetHandle(g)

	g.Cursor = true
	g.Mouse = false
	g.SetManagerFunc(Layout)
	if err := initKeybinding(g); err != nil {
		log.Fatal(err)
	}

	if err := g.MainLoop(); err != nil || err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
