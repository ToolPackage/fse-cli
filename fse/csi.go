package fse

import "fmt"

const (
	cursorUp      = "\033[%dA"
	cursorDown    = "\033[%dB"
	cursorForward = "\033[%dC"
	cursorBack    = "\033[%dD"
)

func moveCursorUp(n int) string {
	return fmt.Sprintf(cursorUp, n)
}

func moveCursorDown(n int) string {
	return fmt.Sprintf(cursorDown, n)
}

func moveCursorForward(n int) string {
	return fmt.Sprintf(cursorForward, n)
}

func moveCursorBack(n int) string {
	return fmt.Sprintf(cursorBack, n)
}
