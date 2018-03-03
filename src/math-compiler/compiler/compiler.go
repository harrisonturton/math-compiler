package main

import (
	"fmt"
	"../parser"
	"../scanner"
	"../token"
	"os"
)

type Compiler struct {
	ast parser.Expr
}

func NewCompiler(ast parser.Expr) *Compiler {
	return &Compiler{ast}
}

func (c *Compiler) compile(expr parser.Expr) string {
	switch expr.(type) {
	case parser.Number:
		num := expr.(parser.Number)
		return c.compileNumber(num)
	case parser.BinaryOp:
		binOp := expr.(parser.BinaryOp)
		return c.compileBinaryOp(binOp)
	}
	return "Compilation finished."
}

func (c *Compiler) compileNumber(num parser.Number) string {
	return fmt.Sprintf("MOV r0, %s", num.Token.Value)
}

func (c *Compiler) compileBinaryOp(binOp parser.BinaryOp) string {
	result := ""
	switch binOp.Left.(type) {
	case parser.Number:
		num := binOp.Left.(parser.Number)
		result += fmt.Sprintf("\nMOV r0, %s", num.Token.Value)
		result += "\nPUSH {r0}"
	case parser.BinaryOp:
		b := binOp.Left.(parser.BinaryOp)
		result += c.compileBinaryOp(b)
	}
	switch binOp.Right.(type) {
	case parser.Number:
		num := binOp.Right.(parser.Number)
		result += fmt.Sprintf("\nMOV r1, %s", num.Token.Value)
		result += "\nPUSH {r1}"
	case parser.BinaryOp:
		b := binOp.Right.(parser.BinaryOp)
		result += c.compileBinaryOp(b)
	}
	result += "\nPOP {r0, r1}"
	switch binOp.Op.Token {
	case token.TOK_ADD:
		result += "\nADD r0, r1"
	case token.TOK_SUB:
		result += "\nSUB r0, r1, r0"
	}
	result += "\nPUSH {r0}"
	return result
}

func parseBytes(input []byte) parser.Expr {
	tokens := scan(input)
	p := parser.NewParser(tokens)
	expr := p.Parse()
	fmt.Println(expr)
	return expr
}

func scan(input []byte) []token.Token {
	s := scanner.NewScanner(input)
	go s.Scan()
	var allTokens []token.Token
	for tok := range s.Tokens {
		allTokens = append(allTokens, tok)
	}
	return allTokens
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("parse -m <input>")
		os.Exit(1)
	}
	if os.Args[1] != "-m" {
		fmt.Println("Need -m")
		os.Exit(1)	
	}
	expr := parseBytes(append([]byte(os.Args[2]), []byte("\r")...))
	/*expr := parser.BinaryOp{
		parser.Number{token.Token{token.TOK_NUMBER, "3"}},
		token.Token{token.TOK_ADD, "+"},
		parser.Number{token.Token{token.TOK_NUMBER, "4"}},
	}*/
	//expr := parser.Number{token.Token{token.TOK_NUMBER, "3"}}
	//expr := parseBytes([]byte("1+2\r"))
	c := NewCompiler(expr)
	fmt.Println(c.compile(expr))
}
