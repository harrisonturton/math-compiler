package parser

import (
	"../token"
	"fmt"
)

type Parser struct {
	tokens    chan token.Token
	current   token.Token
}

func NewParser(tokens chan token.Token) *Parser {
	current := <-tokens
	return &Parser{tokens, current}
}

func (p *Parser) Parse() Expr {
	return p.parseAdditionSequence()
}

// Move the parser one token forwards
// Return new token
func (p *Parser) next() token.Token {
	if p.current.Token == token.TOK_EOF {
		return p.current
	}
	tok := p.current
	p.current = <-p.tokens
	return tok
}

// See next token without shifting parser
func (p *Parser) peek() token.Token {
	return p.current
}

// Move to next token, ignore current one
func (p *Parser) ignore() {
	p.next()
}

// sub => num (- num)* | add
//

/*
P ---> E '$'
E ---> T {('+'|'-') T}
T ---> S {('*'|'/') S}
S ---> F '^' S | F
F ---> '(' E ')' | char
*/

func (p *Parser) parseAdditionSequence() Expr {
	left := p.parseNumber()
	for p.peek().Token == token.TOK_SUB || p.peek().Token == token.TOK_ADD {
		op := p.next()
		right := p.parseNumber()
		left = BinaryOp{left, op, right}
		fmt.Println(fmt.Sprintf("lookahead: %s", p.peek().Token))

		// parse num
		if p.peek().Token != token.TOK_SUB || p.peek().Token != token.TOK_ADD {

		}
	}
	return left
}

func (p *Parser) parseSub() Expr {
	left := p.parseNumber()
	if p.peek().Token != token.TOK_SUB {
		return left
	}
	op := p.next()
	right := p.parseNumber()
	return BinaryOp{left, op, right}
}

func (p *Parser) parseNumber() Expr {
	if p.peek().Token == token.TOK_LPAREN {
		p.ignore() // Ignore starting TOK_LPAREN
		expr := p.parseSub()
		p.ignore() // Ignore closing TOK_RPAREN
		return expr
	}
	tok := p.next()
	return Number{tok}
}
