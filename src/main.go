package main

import (
	"math-compiler/src/compiler"
	"math-compiler/src/parser"
	"math-compiler/src/token"
	"math-compiler/src/scanner"
	"fmt"
	"io/ioutil"
	"os"
)

const errorMsg = `No input specified. Please use either:
	1. <filepath>
	2. -m <input>
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(errorMsg)
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		parseFile(os.Args[1])
		os.Exit(0)
	}
	if len(os.Args) != 3 {
		fmt.Println(errorMsg)
		os.Exit(1)
	}
	if os.Args[1] == "-m" || os.Args[1] == "--manual" {
		parseMessage(os.Args[2])
	}
}

func parseFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	expr := parseBytes(bytes)
	compiled := compile(expr)
	fmt.Println(compiled)
}

func parseMessage(message string) {
	input := append([]byte(message), '\r')
	expr := parseBytes(input)
	compiled := compile(expr)
	fmt.Println(compiled)
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

func compile(expr parser.Expr) string {
	return compiler.Compile(expr)
}
