package ui

import (
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/ToolPackage/fse-cli/fse/version"
	"github.com/jroimartin/gocui"
)

var (
	notify *notificator.Notificator
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	notify = notificator.New(notificator.Options{
		AppName: version.Name,
	})

	if v, err := g.SetView("logo", 0, 0, maxX, 7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Wrap = false
		v.Frame = false

		_, _ = fmt.Fprint(v, version.ColorLogo())
	}

	if v, err := g.SetView("log", -1, 7, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = false
	}

	return nil
}
