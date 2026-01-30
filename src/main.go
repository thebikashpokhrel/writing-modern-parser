package main

import (
	"fmt"
	"os"
	"writing-modern-parser/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/code00.lang")
	source := string(bytes)

	fmt.Printf("Code: %s\n", source)

	tokens := lexer.Tokenize(source)

	for _, token := range tokens {
		token.Debug()
	}
}
