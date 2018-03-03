package parser

import (
	"../token"
)

// Entry point for the parser
func (p *Parser) parseExpr() Expr {
	return p.parseAddTerm()
}

func (p *Parser) parseAddTerm() Expr {
	left := p.parseMulTerm()
	if !isAddition(p.peek()) || isExpressionEnd(p.peek()) {
		return left
	}
	for isAddition(p.peek()) {
		op := p.next()
		right := p.parseMulTerm()
		left = BinaryOp{left, op, right}
	}
	return left
}

func (p *Parser) parseMulTerm() Expr {
	left := p.parseOrdTerm()
	if !isMul(p.peek()) || isExpressionEnd(p.peek()) {
		return left
	}
	for isMul(p.peek()) {
		op := p.next()
		right := p.parseOrdTerm()
		left = BinaryOp{left, op, right}
	}
	return left
}

func (p *Parser) parseOrdTerm() Expr {
	left := p.parseNumber()
	if p.peek().Token != token.TOK_ORD || isExpressionEnd(p.peek()) {
		return left
	}
	for p.peek().Token == token.TOK_ORD {
		op := p.next()
		right := p.parseNumber()
		left = BinaryOp{left, op, right}
	}
	return left
}

func (p *Parser) parseNumber() Expr {
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
		expr := p.parseAddTerm()
		p.ignore() // Ignore closing TOK_RPAREN
		if neg {
			return UnaryOp{token.Token{token.TOK_SUB, "-"}, expr}
		}
		return expr
	}
	tok := p.next()
	if neg {
		return UnaryOp{token.Token{token.TOK_SUB, "-"}, Number{tok}}
	}
	return Number{tok}
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
