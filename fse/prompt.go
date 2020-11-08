package fse

import (
	"bufio"
	"github.com/eiannone/keyboard"
	"os"
	"strings"
)

const prompt = "> "

var input strings.Builder
var output = bufio.NewWriter(os.Stdout)
var inputHistory = NewHistory()

func InputHandler(char rune, key keyboard.Key) error {
	if key == keyboard.KeyEsc {
		return QuitError
	}

	switch {
	case char != 0:
		inputHistory.ResetCurrent()
		OutputRune(char)
		input.WriteRune(char)
	case key == keyboard.KeySpace:
		OutputRune(' ')
		input.WriteRune(' ')
	case key == keyboard.KeyEnter:
		line := strings.Trim(input.String(), " ")
		if len(line) > 0 {
			input.Reset()
			inputHistory.Add(line)
			inputHistory.ResetCurrent()
			OutputRune('\n')
			RenderSuggestion(line)
			OutputString("text: " + line)
			NewLine()
		}

	case key == keyboard.KeyArrowUp:
		// choose history
		if line := inputHistory.Prev(); len(line) > 0 {
			EraseLine()
			OutputString(line)
			input.Reset()
			input.WriteString(line)
		}
	case key == keyboard.KeyArrowDown:
		// choose history
		var line string
		if line = inputHistory.Next(); len(line) == 0 {
			// recover input
			inputHistory.ResetCurrent()
		}
		EraseLine()
		OutputString(line)
		input.Reset()
		input.WriteString(line)
	case key == keyboard.KeyBackspace:
		if Back() {
			inputHistory.ResetCurrent()
		}
	}

	return nil
}

func NewLine() {
	OutputRune('\n')
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	OutputString(path)
	OutputString(prompt)
}

func OutputRune(c rune) {
	_, _ = output.WriteRune(c)
	_ = output.Flush()
}

func OutputString(s string) {
	_, _ = output.WriteString(s)
	_ = output.Flush()
}

func EraseLine() {
	sz := input.Len()
	if sz > 0 {
		input.Reset()
		_, _ = output.WriteString(moveCursorBack(sz))
		_, _ = output.WriteString(strings.Repeat(" ", sz))
		_, _ = output.WriteString(moveCursorBack(sz))
		_ = output.Flush()
	}
}

func Back() bool {
	sz := input.Len()
	if sz > 0 {
		_, _ = output.WriteString(moveCursorBack(1))
		_, _ = output.WriteString(" ")
		_, _ = output.WriteString(moveCursorBack(sz))
		return true
	}
	return false
}

func Forward() {

}
