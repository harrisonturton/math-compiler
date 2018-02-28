package main

import (
	"./parser"
	"./scanner"
	"./token"
	"fmt"
	"io/ioutil"
	"os"
)

func readFile() []byte {
	if len(os.Args) < 2 {
		fmt.Println("Filepath required.")
		os.Exit(1)
	}
	bytes, _ := ioutil.ReadFile(os.Args[1])
	return bytes
}

func scan(input []byte) (*scanner.Scanner, []token.Token) {
	s := scanner.NewScanner(input)
	go s.Scan()
	var allTokens []token.Token
	for tok := <-s.Tokens; tok.Token != token.TOK_EOF; tok = <-s.Tokens {
		allTokens = append(allTokens, tok)
	}
	return s, allTokens
}

func main() {
	input := readFile()
	fmt.Println(string(input))
	_, tokens := scan(input)
	p := parser.NewParser(tokens)
	fmt.Println(p.Parse())
}
