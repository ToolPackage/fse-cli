package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

var (
	inputHistory = NewHistory()
)

func SimpleEditorFunc(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyTab:
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
	case key == gocui.KeyEnter:
		inputHistory.ResetCurrent()
		input, _ := UI.ReadInput()
		inputHistory.Add(input)
		_ = UI.WriteOutput(input)
	case key == gocui.KeyArrowUp:
		// choose history
		if line := inputHistory.Prev(); len(line) > 0 {
			UI.ClearInput()
			_, _ = fmt.Fprint(v, line)
			_ = v.SetCursor(EditorPromptSz+len(line), 0)
		}
	case key == gocui.KeyArrowDown:
		// choose history
		if line := inputHistory.Next(); len(line) > 0 {
			UI.ClearInput()
			_, _ = fmt.Fprint(v, line)
			_ = v.SetCursor(EditorPromptSz+len(line), 0)
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
		_ = v.SetCursor(EditorPromptSz, 0)
		_ = v.SetOrigin(0, 0)
	case key == gocui.KeyCtrlK:
		inputHistory.ResetCurrent()
		// clear input
		UI.ClearInput()
	case key == gocui.KeyCtrlE:
		// move to the end of the line
		_ = v.SetCursor(len(v.Buffer())-1, 0)
	}
}
