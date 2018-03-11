package main

import (
	"github.com/gopherjs/gopherjs/js"
	"math-compiler/src/compiler"
	"math-compiler/src/parser"
	"math-compiler/src/token"
	"math-compiler/src/scanner"
	"fmt"
	"errors"
)

func main() {
	js.Global.Set("Compile", CompileString)
	fmt.Println(CompileString("1+2"))
}

func CompileString(message string) string {
	input := append([]byte(message), '\r')
	expr, err := parseBytes(input)
	if err != nil {
		return err.Error()
	}
	return compiler.Compile(expr)
}

func parseBytes(bytes []byte) (parser.Expr, error) {
	tokens, err := scan(bytes)
	if err != nil {
		return nil, err
	}
	p := parser.NewParser(tokens)
	return p.Parse()
}

func scan(input []byte) ([]token.Token, error) {
	s := scanner.NewScanner(input)
	go s.Scan()
	var allTokens []token.Token
	for tok := range s.Tokens {
		if tok.Token == token.TOK_ERR {
			return nil, errors.New("SYNTAX ERROR: " + tok.Value)
		}
		allTokens = append(allTokens, tok)
	}
	return allTokens, nil
}
