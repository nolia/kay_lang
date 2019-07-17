package lexer

import (
	"fmt"
	"regexp"
	"unicode"
)

var f_debug = false

func log(pattern string, args ...interface{}) {
	if !f_debug {
		return
	}
	fmt.Printf(pattern, args...)
}

func Tokenize(s string) []Token {
	res := []Token{}

	buff := []rune{}

	runes := []rune(s)
	i := 0
	for i < len(runes) {
		r := runes[i]

		switch {
		case unicode.IsSpace(r):
			res, buff = pushToken(res, buff)

		case isOperator(r):
			res, buff = pushToken(res, buff)
			res = append(res, Token{Type: OPERATOR, Value: string(r)})

		case r == '(':
			res, buff = pushToken(res, buff)
			res = append(res, Token{Type: OPEN_PAR, Value: string(r)})

		case r == ')':
			res, buff = pushToken(res, buff)
			res = append(res, Token{Type: CLOSE_PAR, Value: string(r)})

		default:
			buff = append(buff, r)
		}

		i++
	}

	// Append last buffer.
	res, _ = pushToken(res, buff)

	log("%q => %v\n", s, res)

	return res
}

func isOperator(r rune) bool {
	switch r {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}

func pushToken(tokens []Token, buff []rune) ([]Token, []rune) {
	if len(buff) == 0 {
		return tokens, []rune{}
	}

	const (
		regexNumber   = "^\\d*$"
		regexOperator = "[-=*/]"
		regexIdent    = "[a-zA-Z][a-zA-Z0-9_]*"
	)

	s := string(buff)
	matches := func(pattern string) bool {
		input := s
		ok, err := regexp.MatchString(pattern, input)
		return ok && err == nil
	}

	if matches(regexNumber) {
		return append(tokens, Token{Type: NUMBER, Value: s}), []rune{}
	}

	if matches(regexOperator) {
		return append(tokens, Token{Type: OPERATOR, Value: s}), []rune{}
	}

	if matches(regexIdent) {
		// Check it is a keyword.

		return append(tokens, Token{Type: IDENTIFIER, Value: s}), []rune{}
	}

	return tokens, []rune{}
}
