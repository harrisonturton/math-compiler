package main

import (
	"math-compiler/parser"
	"math-compiler/scanner"
	"math-compiler/token"
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
	parseBytes(bytes)
}

func parseMessage(message string) {
	input := append([]byte(message), []byte("\r")...)
	parseBytes(input)
}

func parseBytes(input []byte) {
	tokens := scan(input)
	p := parser.NewParser(tokens)
	fmt.Println(p.Parse())
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
