package scanner

import (
	"math-compiler/src/token"
	"strings"
)

func startState(s *Scanner) stateFn {
	ch := s.next()
	switch {
	case isWhitespace(ch):
		s.ignore()
		return startState
	case ch == eof:
		s.emit(token.TOK_EOF)
		return nil
	case isBracket(ch):
		s.backup()
		return scanBracket
	case isNumber(ch):
		s.backup()
		return scanNumber
	case isOperator(ch):
		s.backup()
		return scanOperator
	default:
		s.errorf("Unknown character [%s], rune [%d]", string(ch), int(ch))
	}
	s.emit(token.TOK_EOF)
	return nil
}

func scanNumber(s *Scanner) stateFn {
	signs := "+-"
	numbers := "0123456789"

	s.accept(signs)
	s.acceptRun(numbers)
	s.accept(".")
	s.acceptRun(numbers)

	s.emit(token.TOK_NUMBER)
	return startState
}

func scanOperator(s *Scanner) stateFn {
	ch := s.next()
	switch ch {
	case '+':
		s.emit(token.TOK_ADD)
		return startState
	case '-':
		s.emit(token.TOK_SUB)
		return startState
	case '*':
		s.emit(token.TOK_MUL)
		return startState
	case '/':
		s.emit(token.TOK_DIV)
		return startState
	case '^':
		s.emit(token.TOK_ORD)
		return startState
	default:
		s.errorf("Unknown operator [%s]", string(ch))
	}
	return startState
}

func scanBracket(s *Scanner) stateFn {
	ch := s.next()
	switch ch {
	case '(':
		s.emit(token.TOK_LPAREN)
		return startState
	case ')':
		s.emit(token.TOK_RPAREN)
		return startState
	default:
		s.errorf("Unknown operator [%s]", string(ch))
	}
	return startState
}

func isBracket(ch rune) bool {
	brackets := "()"
	return strings.IndexRune(brackets, ch) >= 0
}

func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isOpenParen(ch rune) bool {
	return ch == '('
}

func isOperator(ch rune) bool {
	operators := "+-*/^"
	return strings.IndexRune(operators, ch) >= 0
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\n' || ch == '\t' || ch == '\v' || ch == '\f' || ch == '\r'
}
