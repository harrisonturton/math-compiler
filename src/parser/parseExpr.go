package parser

import (
	"math-compiler/src/token"
	"errors"
	"fmt"
)

// Entry point for the parser
func (p *Parser) parseExpr() (Expr, error) {
	return p.parseAddTerm()
}

func (p *Parser) parseAddTerm() (Expr, error) {
	left, err := p.parseMulTerm()
	if err != nil {
		return nil, err
	}
	if !isAddition(p.peek()) || isExpressionEnd(p.peek()) {
		return left, nil
	}
	for isAddition(p.peek()) {
		op := p.next()
		if p.peek().Token == token.TOK_EOF {
			return nil, errors.New("ERROR: Unexpected EOF. Have you forgotten a number?")
		}
		right, err := p.parseMulTerm()
		if err != nil {
			return nil, err
		}
		left = BinaryOp{left, op, right}
	}
	return left, nil
}

func (p *Parser) parseMulTerm() (Expr, error) {
	left, err := p.parseOrdTerm()
	if err != nil {
		return nil, err
	}
	if !isMul(p.peek()) || isExpressionEnd(p.peek()) {
		return left, nil
	}
	for isMul(p.peek()) {
		op := p.next()
		if p.peek().Token == token.TOK_EOF {
			return nil, errors.New("ERROR: Unexpected EOF. Have you forgotten a number?")
		}
		right, err := p.parseOrdTerm()
		if err != nil {
			return nil, err
		}
		left = BinaryOp{left, op, right}
	}
	return left, nil
}

func (p *Parser) parseOrdTerm() (Expr, error) {
	left, err := p.parseNumber()
	if err != nil {
		return nil, err
	}
	if p.peek().Token != token.TOK_ORD || isExpressionEnd(p.peek()) {
		return left, nil
	}
	for p.peek().Token == token.TOK_ORD {
		op := p.next()
		if p.peek().Token == token.TOK_EOF {
			return nil, errors.New("ERROR: Unexpected EOF. Have you forgotten a number?")
		}
		right, err := p.parseNumber()
		if err != nil {
			return nil, err
		}
		left = BinaryOp{left, op, right}
	}
	return left, nil
}

func (p *Parser) parseNumber() (Expr, error) {
	neg := false
	if p.peek().Token == token.TOK_SUB {
		neg = true
		p.ignore()
	}
	if p.peek().Token == token.TOK_ADD {
		p.ignore()
	}
	if p.peek().Token == token.TOK_LPAREN {
		p.ignore() // Ignore starting TOK_LPAREN
		expr, err := p.parseAddTerm()
		if err != nil {
			return nil, err
		}
		p.ignore() // Ignore closing TOK_RPAREN
		if neg {
			return UnaryOp{token.Token{token.TOK_SUB, "-"}, expr}, nil
		}
		return expr, nil
	}
	tok := p.next()
	if tok.Token != token.TOK_NUMBER {
		return nil, errors.New(fmt.Sprintf("Expected number, recieved %s", tok))
	}
	if neg {
		return UnaryOp{token.Token{token.TOK_SUB, "-"}, Number{tok}}, nil
	}
	return Number{tok}, nil
}

func isExpressionEnd(tok token.Token) bool {
	return isEOF(tok) || isCloseParen(tok)
}

func isAddition(tok token.Token) bool {
	return tok.Token == token.TOK_SUB || tok.Token == token.TOK_ADD
}

func isMul(tok token.Token) bool {
	return tok.Token == token.TOK_MUL || tok.Token == token.TOK_DIV
}

func isEOF(tok token.Token) bool {
	return tok.Token == token.TOK_EOF
}

func isCloseParen(tok token.Token) bool {
	return tok.Token == token.TOK_RPAREN
}
