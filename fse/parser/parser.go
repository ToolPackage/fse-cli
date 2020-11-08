package parser

import (
	"strings"
)

func Parse(input string) *TokenStream {
	tokens := strings.Split(input, " ")
	return &TokenStream{next: 0, tokens: tokens}
}

type TokenStream struct {
	next   int
	tokens []string
}

func (ts *TokenStream) Next() (string, bool) {
	for {
		if ts.next >= len(ts.tokens) {
			return "", false
		}
		token := strings.Trim(ts.tokens[ts.next], "\t\r\n ")
		ts.next++
		if len(token) > 0 {
			return token, true
		}
	}
}
