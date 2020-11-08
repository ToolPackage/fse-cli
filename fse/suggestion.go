package fse

import (
	"strconv"
	"strings"
)

var commands = []string{
	"login %address %token",
	"logout",
	"cd %path",
	"ls -r",
	"upload %path",
	"get %path",
	"delete %path",
}

type SuggestionPattern struct {
	text          string
	isPlaceholder bool
}

type Suggestion []SuggestionPattern

var suggestions []Suggestion

func init() {
	suggestions = make([]Suggestion, 0)
	for _, command := range commands {
		sp := make(Suggestion, 0)
		for _, pattern := range strings.Split(command, " ") {
			p := SuggestionPattern{}
			if strings.HasPrefix(pattern, "%") {
				p.text = pattern[1:]
				p.isPlaceholder = true
			} else {
				p.text = pattern
			}
			sp = append(sp, p)
		}
		suggestions = append(suggestions, sp)
	}
}

func RenderSuggestion(input string) {
	tokens := strings.Split(input, " ")
	s := filterSuggestions(tokens)
	if len(s) == 0 {
		return
	}

	deleteToEnd()
	x, y := mainView.Cursor()
	mainView.Overwrite = true

	//outputString(fmt.Sprint(tokens))
	a := s[0]
	for idx := len(tokens); idx < len(a); idx++ {
		pattern := a[idx]
		outputRune(' ')
		outputString(pattern.text + strconv.Itoa(idx))
	}

	_ = mainView.SetCursor(x, y)
}

func filterSuggestions(tokens []string) []Suggestion {
	s := make([]Suggestion, 0)
	for _, suggestion := range suggestions {
		if matchSuggestion(tokens, suggestion) {
			s = append(s, suggestion)
		}
	}
	return s
}

func matchSuggestion(tokens []string, suggestion Suggestion) bool {
	for idx := 0; idx < len(suggestion) && idx < len(tokens); idx++ {
		if !matchSuggestionPattern(tokens[idx], suggestion[idx]) {
			return false
		}
	}
	return true
}

func matchSuggestionPattern(token string, pattern SuggestionPattern) bool {
	if pattern.isPlaceholder {
		return true
	}
	return strings.HasPrefix(pattern.text, token)
}
