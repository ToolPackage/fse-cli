package fse

import (
	"fmt"
	"github.com/ToolPackage/fse-cli/fse/version"
	"github.com/jroimartin/gocui"
	"log"
	"strings"
)

var ui *gocui.Gui
var mainView *gocui.View

const (
	MainView       = "main"
	EditorPrompt   = "> "
	EditorPromptSz = len(EditorPrompt)
)

func Run() {
	var err error
	ui, err = gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer ui.Close()

	ui.Cursor = true
	ui.Mouse = false
	ui.SetManagerFunc(Layout)

	if err := initKeybinding(ui); err != nil {
		log.Fatal(err)
	}

	if err := ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("logo", -1, -1, maxX, 1); err != nil {
		v.Wrap = false
		_, _ = fmt.Fprint(v, version.SimpleLogo())
	}

	if v, err := g.SetView(MainView, -1, 1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = false
		v.Autoscroll = true
		v.Frame = false
		v.Editable = true
		v.Editor = gocui.EditorFunc(simpleEditorFunc)
		_, _ = g.SetCurrentView(MainView)
		mainView = v

		printPrompt()
	}

	return nil
}

func initKeybinding(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	return nil
}

func quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func getInput() string {
	x, y := mainView.Cursor()
	line, _ := mainView.Line(y)
	return strings.Trim(line[EditorPromptSz:x], "\n ")
}

func clearInput() {
	_, y := mainView.Cursor()
	line, _ := mainView.Line(y)
	setCursorToEnd()
	for i := len(line) - EditorPromptSz; i > 0; i-- {
		mainView.EditDelete(true)
	}
}

func deleteToEnd() {
	x, y := mainView.Cursor()
	line, _ := mainView.Line(y)
	n := len(line) - x
	for i := 0; i < n; i++ {
		mainView.EditDelete(false)
	}
	_ = mainView.SetCursor(x, y)
}

func newLine() {
	mainView.EditNewLine()
}

func printPrompt() {
	outputString(EditorPrompt)
}

func outputRune(r rune) {
	mainView.EditWrite(r)
}

func outputString(text string) {
	for _, r := range text {
		mainView.EditWrite(r)
	}
}

func setCursorToEnd() {
	lines := mainView.BufferLines()
	y := len(lines) - 1
	x := len(lines[y])
	_ = mainView.SetCursor(x, y)
}
