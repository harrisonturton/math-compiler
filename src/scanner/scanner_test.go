package scanner

import (
	"math-compiler/src/token"
	"fmt"
	"testing"
)

type testCase struct {
	input    []byte
	expected []token.Token
}

// Used to easily test multiple different inputs
var testCases = [...]testCase{
	input1,
	input2,
	input3,
}

var input1 = testCase{
	input: []byte("1+2"),
	expected: []token.Token{
		{token.TOK_NUMBER, "1"},
		{token.TOK_ADD, "+"},
		{token.TOK_NUMBER, "2"},
	},
}

var input2 = testCase{
	input: []byte("1+-2"),
	expected: []token.Token{
		{token.TOK_NUMBER, "1"},
		{token.TOK_ADD, "+"},
		{token.TOK_SUB, "-"},
		{token.TOK_NUMBER, "2"},
	},
}

var input3 = testCase{
	input: []byte("+1.123+-2.0"),
	expected: []token.Token{
		{token.TOK_ADD, "+"},
		{token.TOK_NUMBER, "1.123"},
		{token.TOK_ADD, "+"},
		{token.TOK_SUB, "-"},
		{token.TOK_NUMBER, "2.0"},
	},
}

func TestScanner(t *testing.T) {
	for _, input := range testCases {
		testScan(input, t)
	}
}

func testScan(input testCase, t *testing.T) {
	input.input = append(input.input, []byte("\r")...) // Required to be valid utf-8 file
	result := scanAllTokens(input.input)
	if !testEq(result, input.expected) {
		t.Error(fmt.Sprintf("Expected %s but recieved %s", input.expected, result))
	}
}

func scanAllTokens(input []byte) []token.Token {
	s := NewScanner(input)
	go s.Scan()
	var allTokens []token.Token
	for tok := <-s.Tokens; tok.Token != token.TOK_EOF; tok = <-s.Tokens {
		allTokens = append(allTokens, tok)
	}
	return allTokens
}

func testEq(a, b []token.Token) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
