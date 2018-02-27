package scanner

import (
	"testing"
	"../token"
)

func scan(input []byte) (*Scanner, chan token.Token) {
	s := NewScanner(input)
	go s.Scan()
	return s, s.Tokens
}

func TestScanner(t *testing.T) {
	_, result := scan([]byte("1+2"))	
	expected := []token.Token{
		{token.TOK_NUMBER, "1"},
		{token.TOK_ADD, "+"},
		{token.TOK_NUMBER, "2"},
	}
	if result != expected {
		t.Error(fmt.Sprintf("Expected %s\nRecieved %s", expected, result))
	}
}
