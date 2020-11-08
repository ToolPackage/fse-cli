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
	case key == keyboard.KeyEnter:
		line := input.String()
		if len(line) > 0 {
			input.Reset()
			inputHistory.Add(line)
			inputHistory.ResetCurrent()
			OutputRune('\n')
			OutputString("text: " + line)
			NewLine()
		}
	case key == keyboard.KeyArrowUp:
		if line := inputHistory.Prev(); len(line) > 0 {
			OutputString(eraseLine)
			OutputString(line)
			input.WriteString(line)
		}
	case key == keyboard.KeyArrowDown:
		if line := inputHistory.Next(); len(line) > 0 {
			OutputString(eraseLine)
			OutputString(eraseLine)
			OutputString(line)
			input.WriteString(line)
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
	for i := 0; i < sz; i++ {
		_, _ = output.WriteString(cursorBack)
	}

}
