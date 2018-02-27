package parser

import (
	"../token"
	//"fmt"
)

type Parser struct {
	tokens []token.Token
	pos    int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) Parse() Expr {
	return p.parseExpr()
}

// Move the parser one token forwards
// Return new token
func (p *Parser) next() token.Token {
	if p.pos >= len(p.tokens) || p.tokens[p.pos].Token == token.TOK_EOF {
		return token.Token{token.TOK_EOF, ""}
	}
	tok := p.tokens[p.pos]
	p.pos += 1
	return tok
}

// See next token without shifting parser
func (p *Parser) peek() token.Token {
	tok := p.next()
	p.backup()
	return tok
}

// Step back one token
func (p *Parser) backup() {
	p.pos -= 1
}

// Move to next token, ignore current one
func (p *Parser) ignore() {
	p.next()
}
