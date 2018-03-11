package main

import (
	"math-compiler/src/parser"
	"math-compiler/src/scanner"
	"math-compiler/src/token"
	"fmt"
	"io/ioutil"
	"os"
	"errors"
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
	parseBytes(bytes)
}

func parseMessage(message string) {
	input := append([]byte(message), []byte("\r")...)
	parseBytes(input)
}

func parseBytes(input []byte) {
	tokens, err := scan(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	p := parser.NewParser(tokens)
	expr, err := p.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(expr)
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
