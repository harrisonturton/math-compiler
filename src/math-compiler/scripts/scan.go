package main

import (
	"math-compiler/scanner"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("scan <filename> or scan -m <input>")
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		scanFile(os.Args[1])
	}
	if os.Args[1] == "-m" || os.Args[1] == "--manual" {
		scanMessage(os.Args[2])
	}
}

func scanFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	scanBytes(bytes)
}

func scanMessage(message string) {
	input := append([]byte(message), []byte("\r")...)
	scanBytes(input)
}

func scanBytes(input []byte) {
	s := scanner.NewScanner(input)
	go s.Scan()
	for tok := range s.Tokens {
		fmt.Println(fmt.Sprintf("%s %s", tok.Token, tok.Value))
	}
}
