package util

import (
	"fmt"
	"unicode"
)

type TokenType int32

const (
	LPAREN TokenType = iota
	RPAREN
	LAMBDA
	TERM
	EOF
	DOT
)

func tokenTypeToString(t TokenType) string {
	if t == LPAREN {
		return "("
	} else if t == RPAREN {
		return ")"
	} else if t == LAMBDA {
		return "λ"
	} else if t == TERM {
		return "TERM"
	} else if t == EOF {
		return "EOF"
	} else if t == DOT {
		return "."
	} else {
		return "UNEXPECTED TOKEN FOR CONVERSION"
	}
}

type Token struct {
	tokenType TokenType
	termValue string // Non-empty if token is TERM
}

func newToken(tokenType TokenType, termValue string) Token {
	if tokenType != TERM {
		return Token{tokenType, ""}
	} else {
		return Token{tokenType, termValue}
	}
}

func Tokenize(input string) []Token {
	var result []Token
	currentTerm := ""
	if len(input) == 0 {
		return result
	}
	for _, ch := range input {
		nextToken := Token{EOF, ""}
		if ch == 'λ' {
			nextToken = newToken(LAMBDA, "")
		} else if ch == '(' {
			nextToken = newToken(LPAREN, "")
		} else if ch == '.' {
			nextToken = newToken(DOT, "")
		} else if unicode.IsSpace(ch) {
			// Do nothing
		} else if ch == ')' {
			nextToken = newToken(RPAREN, "")
		} else {
			currentTerm = currentTerm + string(ch)
			continue
		}

		if len(currentTerm) > 0 {
			result = append(result, newToken(TERM, currentTerm))
			currentTerm = ""
		}
		
		if nextToken.tokenType != EOF {
			result = append(result, nextToken)
		}
	}
	return result
}

func DisplayLexedInput(t []Token) {
	fmt.Print("\n")
	for _, item := range t {
		if item.tokenType == LPAREN {
			fmt.Print("(")
		} else if item.tokenType == RPAREN {
			fmt.Print(")")
		} else if item.tokenType == LAMBDA {
			fmt.Print("λ")
		} else if item.tokenType == TERM {
			fmt.Print(item.termValue)
		} else if item.tokenType == EOF {
			fmt.Print("EOF")
		} else if item.tokenType == DOT {
			fmt.Print(".")
		} else {
			fmt.Print("UNEXPECTED TOKEN TO DISPLAY")
		}
	}
	fmt.Print("\n")
}
