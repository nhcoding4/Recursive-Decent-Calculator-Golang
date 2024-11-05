package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	instructions()

	for {
		input := takeInput("> ")
		if len(input) == 0 {
			os.Exit(1)
		}

		tokens := lexInput(input)
		if tokens == nil {
			continue
		}

		parse(tokens)
	}
}

func instructions() {
	fmt.Println("This calculator supports: (), +, -, * and / operations.")
	fmt.Println("Enter a blank line to exit.")
}

func lexInput(input string) []Token {
	lexer := NewLexer(input)
	tokens := lexer.CreateTokens()
	errors := lexer.LexingErrors()

	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}

		return nil
	}

	return tokens
}

func parse(tokens []Token) {
	parsedTokens := NewParser(tokens).ParseExpr()
	fmt.Println(parsedTokens.Eval())
}

func takeInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(input)
}
