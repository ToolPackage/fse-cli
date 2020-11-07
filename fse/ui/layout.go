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

	if v, err := g.SetView("logo", -1, -1, maxX, 7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Wrap = false
		v.Frame = true

		_, _ = fmt.Fprint(v, version.ColorLogo())
	}

	if v, err := g.SetView("output", -1, 7, maxX/2, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = false
		v.Frame = false
	}

	if v, err := g.SetView("status", maxX/2, 7, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = false
		v.Frame = false
	}

	if v, err := g.SetView("editor", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.BgColor = gocui.ColorRed
		g.SetCurrentView("editor")
	}

	return nil
}
