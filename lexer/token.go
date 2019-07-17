package lexer

type TokenType string

const (
	NUMBER     TokenType = "number"
	OPERATOR   TokenType = "operator"
	IDENTIFIER TokenType = "ident"
	OPEN_PAR   TokenType = "open_paren"
	CLOSE_PAR  TokenType = "open_paren"
	OTHER      TokenType = "other"
)

type Token struct {
	Type  TokenType
	Value string
}
