package parser

import (
	"math-compiler/src/token"
	"fmt"
	"testing"
)

var input = []token.Token{
	{token.TOK_NUMBER, "3"},
	{token.TOK_ADD, "+"},
	{token.TOK_NUMBER, "4"},
}

func TestParseAll(t *testing.T) {
	for _, input := range testInputs {
		testParse(input, t)
	}
}

func testParse(input testInput, t *testing.T) {
	p := NewParser(input.input)
	result := fmt.Sprintf("%s", p.Parse())
	expected := input.expected
	if result != expected {
		t.Error(fmt.Sprintf("Expected %s\nRecieved %s", expected, result))
	}
}

func makeNumberExpr(num string) Expr {
	return Number{token.Token{token.TOK_NUMBER, num}}
}
