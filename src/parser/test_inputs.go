package parser

import (
	"math-compiler/src/token"
)

type testInput struct {
	input    []token.Token
	expected string
}

var testInputs = [...]testInput{
	inputAdd1,
	inputAdd2,
	inputSub1,
	inputSub2,
	inputSub3,
}

var inputAdd1 = testInput{
	input: []token.Token{
		{token.TOK_NUMBER, "3"},
		{token.TOK_ADD, "+"},
		{token.TOK_NUMBER, "4"},
	},
	expected: "(+ 3 4)",
}

var inputAdd2 = testInput{
	input: []token.Token{
		{token.TOK_NUMBER, "3"},
		{token.TOK_ADD, "+"},
		{token.TOK_NUMBER, "4"},
		{token.TOK_ADD, "+"},
		{token.TOK_NUMBER, "5"},
	},
	expected: "(+ (+ 3 4) 5)",
}

var inputSub1 = testInput{
	input: []token.Token{
		{token.TOK_NUMBER, "3"},
		{token.TOK_SUB, "-"},
		{token.TOK_NUMBER, "4"},
	},
	expected: "(- 3 4)",
}

var inputSub2 = testInput{
	input: []token.Token{
		{token.TOK_NUMBER, "3"},
		{token.TOK_SUB, "-"},
		{token.TOK_NUMBER, "4"},
		{token.TOK_SUB, "-"},
		{token.TOK_NUMBER, "5"},
	},
	expected: "(- (- 3 4) 5)",
}

var inputSub3 = testInput{
	input: []token.Token{
		{token.TOK_NUMBER, "3"},
		{token.TOK_SUB, "-"},
		{token.TOK_LPAREN, "("},
		{token.TOK_NUMBER, "-4"},
		{token.TOK_SUB, "-"},
		{token.TOK_NUMBER, "5"},
		{token.TOK_RPAREN, ")"},
	},
	expected: "(- 3 (- -4 5))",
}
