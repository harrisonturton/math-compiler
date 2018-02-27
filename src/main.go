package main

import (
	"./token"
	"./scanner"
	"./parser"
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

func scan(input []byte) (*scanner.Scanner, chan token.Token) {
	s := scanner.NewScanner(input)
	go s.Scan()
	return s, s.Tokens
}

func main() {
	input := readFile()
	_, tokens := scan(input)
	p := parser.NewParser(tokens)
	fmt.Println(p.Parse())
}
