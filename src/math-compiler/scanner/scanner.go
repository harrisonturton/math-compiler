package scanner

import (
	"math-compiler/token"
	"fmt"
	"strings"
	"unicode/utf8"
)

const eof = rune(-1)

type Scanner struct {
	source []byte

	start int
	pos   int
	width int

	Tokens chan token.Token
}

// Used to represent state transitions
type stateFn func(*Scanner) stateFn

func NewScanner(source []byte) *Scanner {
	return &Scanner{source, 0, 0, 0, make(chan token.Token)}
}

// Lex input by repeatedly executing state functions
func (s *Scanner) Scan() chan token.Token {
	for state := startState; state != nil; {
		state = state(s)
	}
	close(s.Tokens)
	return s.Tokens
}

// Return character at position, shift scanner forwards
func (s *Scanner) next() rune {
	if s.pos >= len(s.source) {
		return eof
	}
	char, width := utf8.DecodeRuneInString(string(s.source[s.pos:]))
	s.width = width
	s.pos += width
	return char
}

// Consume pending input, place token on channel
func (s *Scanner) emit(tokenType token.TokenType) {
	value := string(s.source[s.start:s.pos])
	s.Tokens <- token.Token{tokenType, value}
	s.start = s.pos
}

// Don't shift scanner, but see next character
func (s *Scanner) peek() rune {
	char := s.next()
	s.backup()
	return char
}

func (s *Scanner) peekBack() rune {
	s.backup()
	s.backup()
	char := s.next()
	s.next()
	return char
}

// Step back one rune
func (s *Scanner) backup() {
	s.pos -= s.width
}

func (s *Scanner) ignore() {
	s.start = s.pos
}

// Emit an error and stop lexing
func (s *Scanner) errorf(format string, args ...interface{}) stateFn {
	s.Tokens <- token.Token{
		token.TOK_ERR,
		fmt.Sprintf(format, args...),
	}
	return nil
}

// Consume a single rune if it exists in the string
func (s *Scanner) accept(valid string) bool {
	if strings.IndexRune(valid, s.next()) >= 0 {
		return true
	}
	s.backup()
	return false
}

// Consume series of runes if they exist in the string
func (s *Scanner) acceptRun(valid string) {
	for strings.IndexRune(valid, s.next()) >= 0 {
	}
	s.backup()
}
