package fse

import (
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
	"strings"
	"time"
)

var (
	Client *CommandClient
	Notify *notificator.Notificator
	UI     *Gui
)

type Gui struct {
	g *gocui.Gui
}

func (gui *Gui) SetHandle(g *gocui.Gui) {
	gui.g = g
}

func (gui *Gui) SuccessOutput(text string) error {
	return gui.SendOutput(color.GreenString(text))
}

func (gui *Gui) WarningOutput(text string) error {
	return gui.SendOutput(color.YellowString(text))
}

func (gui *Gui) ErrorOutput(text string) error {
	return gui.SendOutput(color.RedString(text))
}

func (gui *Gui) SendOutput(text string) error {
	v, err := gui.g.View(OutputView)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(v, "[%s] %s\n", getTimestamp(), text)
	return err
}

func getTimestamp() string {
	n := time.Now()
	return fmt.Sprintf("%02d:%02d:%02d.%03d", n.Hour(), n.Minute(), n.Second(), n.Nanosecond()/1e6)
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
