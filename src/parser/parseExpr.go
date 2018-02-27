package parser

import (
	"../token"
	"fmt"
)

/*
 * GRAMMAR
 * =======
 * expr    => mulTerm { +|- mulTerm } $
 * multerm => facTerm { *|/ facTerm }
 * facTerm => num ^ num | num
 * num     => int | ( expr )
 */

// Entry point for the parser
func (p *Parser) parseExpr() Expr {
	return p.parseAddTerm()
}

// 1. Brackets, 2. Orders, 3. Mul, 4. Div, 5. Add, 6. Sub

func (p *Parser) parseOrdTerm() Expr {
	left := p.parseNumber()
	if isEOF(p.peek()) || p.peek().Token != token.TOK_ORD || isCloseParen(p.peek()) {
		return left
	}
	for p.peek().Token == token.TOK_ORD {
		op := p.next()
		right := p.parseNumber()
		left = BinaryOp{left, op, right}
	}
	return left
}

func (p *Parser) parseAddTerm() Expr {
	left := p.parseMulTerm()
	if !isAddition(p.peek()) || isEOF(p.peek()) || isCloseParen(p.peek()) {
		return left
	}
	for isAddition(p.peek()) {
		op := p.next()
		fmt.Println(p.peek())
		right := p.parseMulTerm()
		left = BinaryOp{left, op, right}
	}
	return left
}

func (p *Parser) parseMulTerm() Expr {
	left := p.parseOrdTerm()
	if isEOF(p.peek()) || !isMul(p.peek()) || isCloseParen(p.peek()) {
		return left
	}
	for isMul(p.peek()) {
		op := p.next()
		right := p.parseOrdTerm()
		left = BinaryOp{left, op, right}
	}
	return left
}

func (p *Parser) parseNumber() Expr {
	if p.peek().Token == token.TOK_LPAREN {
		p.ignore() // Ignore starting TOK_LPAREN
		expr := p.parseExpr()
		p.ignore() // Ignore closing TOK_RPAREN
		return expr
	}
	tok := p.next()
	return Number{tok}
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
