package ui

import (
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/jroimartin/gocui"
	"strings"
)

var (
	Notify *notificator.Notificator
	UI     *Gui
)

type Gui struct {
	g *gocui.Gui
}

func (gui *Gui) SetHandle(g *gocui.Gui) {
	gui.g = g
}

func (gui *Gui) WriteOutput(text string) error {
	v, err := gui.g.View(OutputView)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(v, text)
	return err
}

func (gui *Gui) ReadInput() (string, error) {
	v, err := gui.g.View(EditorView)
	if err != nil {
		return "", err
	}
	content := v.Buffer()[EditorPromptSz:]
	gui.ClearInput()
	content = strings.Trim(content, "\r\t\n ")
	return content, nil
}

func (gui *Gui) ClearInput() {
	v, err := gui.g.View(EditorView)
	if err == nil {
		printPrompt(v)
	}
}

func printPrompt(v *gocui.View) {
	v.Clear()
	_, _ = fmt.Fprint(v, EditorPrompt)
	_ = v.SetCursor(EditorPromptSz, 0)
}
