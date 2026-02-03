package main

import (
	"fmt"
	"os"
	"writing-modern-parser/src/lexer"
	"writing-modern-parser/src/parser"

	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("./examples/code04.lang")
	source := string(bytes)

	fmt.Printf("Code: %s\n", source)

	tokens := lexer.Tokenize(source)

	for _, token := range tokens {
		token.Debug()
	}

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
