package main

import (
	"github.com/gopherjs/gopherjs/js"
	"math-compiler/src/compiler"
	"math-compiler/src/parser"
	"math-compiler/src/token"
	"math-compiler/src/scanner"
	"fmt"
)

func main() {
	js.Global.Set("Compile", CompileString)
	fmt.Println(CompileString("1+2"))
}

func CompileString(message string) string {
	input := append([]byte(message), '\r')
	expr := parseBytes(input)
	return compiler.Compile(expr)
}

func parseBytes(bytes []byte) parser.Expr {
	tokens := scan(bytes)
	p := parser.NewParser(tokens)
	return p.Parse()
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
