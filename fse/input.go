package fse

import (
	"github.com/jroimartin/gocui"
)

var (
	inputHistory = NewHistory()
)

func simpleEditorFunc(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyTab:
		// TODO:
	case ch != 0 && mod == 0:
		inputHistory.ResetCurrent()
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		inputHistory.ResetCurrent()
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		cx, _ := v.Cursor()
		if cx > EditorPromptSz {
			inputHistory.ResetCurrent()
			v.EditDelete(true)
		}
	case key == gocui.KeyDelete:
		cx, _ := v.Cursor()
		if cx < len(v.ViewBuffer())-1 {
			inputHistory.ResetCurrent()
			v.EditDelete(false)
		}
	case key == gocui.KeyInsert:
		v.Overwrite = !v.Overwrite
	case key == gocui.KeyArrowUp:
		// choose history
		if line := inputHistory.Prev(); len(line) > 0 {
			clearInput()
			outputString(line)
		}
	case key == gocui.KeyArrowDown:
		// choose history
		clearInput()
		if line := inputHistory.Next(); len(line) > 0 {
			outputString(line)
		} else {
			inputHistory.ResetCurrent()
		}
	case key == gocui.KeyArrowLeft:
		// move cursor
		cx, _ := v.Cursor()
		if cx > EditorPromptSz {
			v.MoveCursor(-1, 0, false)
		}
	case key == gocui.KeyArrowRight:
		// move cursor
		cx, _ := v.Cursor()
		if cx < len(v.ViewBuffer())-1 {
			v.MoveCursor(1, 0, false)
		}
	case key == gocui.KeyCtrlA:
		// move to the start of the line
		x, _ := v.Cursor()
		v.MoveCursor(-x+EditorPromptSz, 0, true)
	case key == gocui.KeyCtrlK:
		// clear input
		inputHistory.ResetCurrent()
		clearInput()
	case key == gocui.KeyCtrlE:
		// move to the end of the line
		x, y := v.Cursor()
		line, _ := v.Line(y)
		v.MoveCursor(len(line)-x, 0, true)
	case key == gocui.KeyEnter:
		// trim line
		line := getInput()
		if len(line) > 0 {
			// move cursor to the end
			setCursorToEnd()
			// update history
			inputHistory.Add(line)
			inputHistory.ResetCurrent()
			// execute
			//Client.Execute(input)
			// print prompt
			newLine()
			printPrompt()
		}
		return
	}
	RenderSuggestion(getInput())
}
