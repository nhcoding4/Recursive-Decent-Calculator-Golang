package main

import (
	"fmt"
	"strconv"
)

const EOF byte = 0

// Splits a string input into tokens. Checks for incorrect input.

type Lexer struct {
	input                                 string
	inputLength, current, peek, openParen int
	ch                                    byte
	errors                                []string
}

// --------------------------------------------------------------------------------------------------------------------

func NewLexer(rawInput string) *Lexer {
	lexer := &Lexer{
		input:       rawInput,
		inputLength: len(rawInput),
		current:     0,
		peek:        0,
	}
	lexer.readChar()

	return lexer
}

// --------------------------------------------------------------------------------------------------------------------
// Lexing functions. Turns raw string data into tokens and finds input errors.
// --------------------------------------------------------------------------------------------------------------------

func (l *Lexer) CreateTokens() []Token {
	tokens := make([]Token, 0)

	for {
		if l.ch == EOF {
			break
		}

		switch l.ch {
		case '+', '-', '/', '*', '(', ')':
			tokens = append(tokens, l.parseOperator())
		default:
			if l.isDigit() {
				tokens = append(tokens, l.parseNumber())
			} else {
				l.parseOther()
			}
		}
	}

	return tokens
}

// --------------------------------------------------------

func (l *Lexer) parseNumber() Token {
	start := l.current

	for {
		if l.ch == EOF || !l.isDigit() {
			break
		}

		l.readChar()
	}

	offset := 0
	if l.ch == EOF {
		offset++
	}

	rawFloat, err := strconv.ParseFloat(l.input[start:l.current+offset], 64)
	if err != nil {
		l.errors = append(l.errors, fmt.Sprintf("Invalid number found at position %v: %v", start, string(l.input[start])))
	}

	return &Number{value: rawFloat}
}

// --------------------------------------------------------

func (l *Lexer) parseOperator() Token {
	l.parensCheck()
	operator := Operator{value: string(l.ch)}
	l.readChar()

	return &operator
}

// --------------------------------------------------------

func (l *Lexer) parseOther() {
	if l.ch != ' ' {
		l.errors = append(l.errors, fmt.Sprintf("Invalid token found at position %v : %v", l.current+1, string(l.ch)))
	}
	l.readChar()
}

// --------------------------------------------------------------------------------------------------------------------
// Helper functions
// --------------------------------------------------------------------------------------------------------------------

func (l *Lexer) isDigit() bool {
	return l.ch >= '0' && l.ch <= '9' || l.ch == '.'
}

// --------------------------------------------------------------------------------------------------------------------

func (l *Lexer) LexingErrors() []string {
	if l.openParen != 0 {
		l.errors = append(l.errors, "Unclosed parenthesis detected")
	}

	return l.errors
}

// --------------------------------------------------------------------------------------------------------------------

func (l *Lexer) parensCheck() {
	switch l.ch {
	case '(':
		l.openParen++
	case ')':
		l.openParen--
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (l *Lexer) readChar() {
	if l.peek >= l.inputLength {
		l.ch = EOF
	} else {
		l.current = l.peek
		l.peek++
		l.ch = l.input[l.current]
	}
}

// --------------------------------------------------------------------------------------------------------------------
