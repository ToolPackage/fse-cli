package fse

import (
	"fmt"
	"github.com/ToolPackage/fse-cli/fse/version"
	"github.com/jroimartin/gocui"
)

const (
	LogoView   = "logo"
	OutputView = "output"
	EditorView = "editor"

	EditorPrompt   = "> "
	EditorPromptSz = len(EditorPrompt)
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView(LogoView, -1, -1, maxX, 7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = false
		v.Frame = true

		_, _ = fmt.Fprint(v, version.ColorLogo())
	}

	if v, err := g.SetView(OutputView, -1, 7, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = false
		v.Frame = false
		v.Autoscroll = true
	}

	if v, err := g.SetView(EditorView, -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editor = gocui.EditorFunc(SimpleEditorFunc)
		v.Editable = true
		v.Frame = false
		v.BgColor = gocui.ColorMagenta
		printPrompt(v)
		if _, err = g.SetCurrentView(EditorView); err != nil {
			return err
		}
	}

	return nil
}
