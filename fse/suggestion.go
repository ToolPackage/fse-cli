package fse

import (
	"fmt"
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
	s := filterSuggestions(input)
	if len(s) == 0 {
		return
	}

	fmt.Println(s)
}

func filterSuggestions(input string) []Suggestion {
	s := make([]Suggestion, 0)
	tokens := strings.Split(input, " ")
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
