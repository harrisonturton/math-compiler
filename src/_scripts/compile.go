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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("parse <filename> or parse -m <input>")
	}
	if len(os.Args) == 2 {
		parseFile(os.Args[1])
	}
	if os.Args[1] == "-m" || os.Args[1] == "--manual" {
		parseMessage(os.Args[2])
	}
}

func parseFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	expr := parseBytes(bytes)
	compile(expr)
}

func compile(expr parser.Expr) {
	generated := compiler.Compile(expr)
	fmt.Println(generated)
}

func parseMessage(message string) {
	input := append([]byte(message), []byte("\r")...)
	expr := parseBytes(input)
	compile(expr)
}

func parseBytes(input []byte) parser.Expr {
	tokens := scan(input)
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
