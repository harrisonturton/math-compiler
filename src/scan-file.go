package main

import (
	"./scanner"
	"./token"
	"fmt"
	"io/ioutil"
	"os"
)

func readFile() []byte {
	if len(os.Args) < 2 {
		fmt.Println("Commands required.")
		os.Exit(1)
	}
	if os.Args[1] == "--input" {
		fmt.Println(os.Args[2])
		return append([]byte(os.Args[2]), []byte("\r")...)
		//return []byte(os.Args[2])
	}
	bytes, _ := ioutil.ReadFile(os.Args[1])
	return bytes
}

func scan(input []byte) []token.Token {
	s := scanner.NewScanner(input)
	go s.Scan()
	var allTokens []token.Token
	for tok := <-s.Tokens; tok.Token != token.TOK_EOF; tok = <-s.Tokens {
		allTokens = append(allTokens, tok)
	}
	return allTokens
}

func main() {
	input := readFile()
	tokens := scan(input)
	for _, tok := range tokens {
		fmt.Println(tok)
	}
}
