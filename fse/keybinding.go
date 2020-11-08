package fse

import "github.com/jroimartin/gocui"

func initKeybinding(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyF1, gocui.ModNone, focusEditorView); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF2, gocui.ModNone, focusOutputView); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	return nil
}

func focusEditorView(g *gocui.Gui, _ *gocui.View) error {
	_, err := g.SetCurrentView(EditorView)
	return err
}

func focusOutputView(g *gocui.Gui, _ *gocui.View) error {
	_, err := g.SetCurrentView(OutputView)
	return err
}

func quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}
