package scanner

import (
	"math-compiler/src/token"
	"testing"
)

func TestNewScanner(t *testing.T) {
	input := []byte("1+2")
	s := NewScanner(input)	
	if len(s.source) != len(input) {
		t.Errorf("Expected scanners.source to be %s, got %s", input, s.source)
	}
	if s.start != 0 {
		t.Errorf("Expected scanner.start to be %s, got %s", 0, s.start)
	}
	if s.pos != 0 {
		t.Errorf("Expected scanner.pos to be %s, got %s", 0, s.pos)
	}
	if s.width != 0 {
		t.Errorf("Expected scanner.width to be %s, got %s", 0, s.width)
	}
}

func TestNext(t *testing.T) {
	input := []byte("1")
	s := NewScanner(input)

	next := s.next()
	if next != '1' {
		t.Errorf("Expected next() to be %s, got %s", '1', next)
	}
	if s.width != 1 {
		t.Errorf("Expected width to be %s, got %s", 1, s.width)
	}

	next = s.next()
	if next != rune(-1) {
		t.Errorf("Expected EOF rune, got %s", next)
	}
}

func TestEmit(t *testing.T) {
	return
	input := []byte("3+2")
	s := NewScanner(input)

	s.next()
	s.emit(token.TOK_NUMBER)
	tok := <-s.Tokens
	expected := token.Token{token.TOK_NUMBER, "3"}
	if tok != expected {
		t.Errorf("Expected token %s, recieved %s", expected, tok)
	}

	s.next()
	s.emit(token.TOK_ADD)
	tok = <-s.Tokens
	expected = token.Token{token.TOK_ADD, "+"}
	if tok != expected {
		t.Errorf("Expected token %s, recieved %s", expected, tok)
	}
}

func TestPeek(t *testing.T) {
	input := []byte("3+2")
	s := NewScanner(input)
	tok := s.peek()
	expected := '3'
	if tok != expected {
		t.Errorf("Expected token %s, recieved %s", expected, tok)
	}

	s.next()
	tok = s.peek()
	expected = '+'
	if tok != expected {
		t.Errorf("Expected token %s, recieved %s", expected, tok)
	}
}

func TestPeekBack(t *testing.T) {
	input := []byte("3+2")
	s := NewScanner(input)

	s.next()
	s.next()
	tok := s.peekBack()
	expected := '3'
	if tok != expected {
		t.Errorf("Expected token %s, recieved %s", expected, tok)
	}

	s.next()
	tok = s.peekBack()
	expected = '+'
	if tok != expected {
		t.Errorf("Expected token %s, recieved %s", expected, tok)
	}
}

func TestBackup(t *testing.T) {
	input := []byte("3")
	s := NewScanner(input)

	s.next()
	s.backup()
	expected := '3'
	tok := s.next()
	if tok != expected {
		t.Errorf("Expected token %s, recieved %s", expected, tok)
	}
}

func TestIgnore(t *testing.T) {
	input := []byte("3+2")
	s := NewScanner(input)

	s.next()
	s.ignore()
	tok := s.next()
	expected := '+'
	if tok != expected {
		t.Errorf("Expected ignored token %s, recieved %s", expected, tok)
	}
}
