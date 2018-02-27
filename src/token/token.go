package token

import "fmt"

type TokenType int

type Token struct {
	Token TokenType
	Value string
}

const (
	TOK_ERR TokenType = iota
	TOK_EOF

	TOK_NUMBER
	TOK_SUB
	TOK_ADD
	TOK_DIV
	TOK_MUL
	TOK_ORD

	TOK_LPAREN
	TOK_RPAREN
)

var tokenDisplay = [...]string{
	"TOK_ERR",
	"TOK_EOF",

	"TOK_NUMBER",
	"TOK_SUB",
	"TOK_ADD",
	"TOK_DIV",
	"TOK_MUL",
	"TOK_ORD",

	"TOK_LPAREN",
	"TOK_RPAREN",
}

func (t TokenType) String() string {
	s := ""
	if 0 <= t && t < TokenType(len(tokenDisplay)) {
		s = tokenDisplay[t]
	} else {
		s = fmt.Sprintf("TokenType(%d)", int(t))
	}
	return s
}

func (t Token) String() string {
	switch t.Token {
	case TOK_EOF:
		return "EOF"
	case TOK_ADD:
		return "+"
	case TOK_SUB:
		return "-"
	}
	if t.Token == TOK_EOF {
		return "EOF"
	}
	return fmt.Sprintf("[%s %s]", t.Token.String(), t.Value)
}
