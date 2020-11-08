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
var inHistory bool

func InputHandler(char rune, key keyboard.Key) error {
	if key == keyboard.KeyEsc {
		return QuitError
	}

	switch {
	case char != 0:
		inputHistory.ResetCurrent()
		OutputRune(char)
		input.WriteRune(char)
		inHistory = false
	case key == keyboard.KeyEnter:
		line := input.String()
		if inHistory {
			line = inputHistory.Current()
		}
		if len(line) > 0 {
			input.Reset()
			inputHistory.Add(line)
			inputHistory.ResetCurrent()
			OutputRune('\n')
			OutputString("text: " + line)
			NewLine()
		}
		inHistory = false
	case key == keyboard.KeyArrowUp:
		// choose history
		if line := inputHistory.Prev(); len(line) > 0 {
			EraseLine()
			OutputString(line)
			inHistory = true
		}
	case key == keyboard.KeyArrowDown:
		// choose history
		EraseLine()
		if line := inputHistory.Next(); len(line) > 0 {
			OutputString(line)
			inHistory = true
		} else {
			inputHistory.ResetCurrent()
			OutputString(input.String())
			inHistory = false
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
	var sz int
	if inHistory {
		sz = len(inputHistory.Current())
	} else {
		sz = input.Len()
		input.Reset()
	}
	for i := 0; i < sz; i++ {
		_, _ = output.WriteString(cursorBack)
	}
	for i := 0; i < sz; i++ {
		_, _ = output.WriteRune(' ')
	}
	for i := 0; i < sz; i++ {
		_, _ = output.WriteString(cursorBack)
	}
	_ = output.Flush()
}
