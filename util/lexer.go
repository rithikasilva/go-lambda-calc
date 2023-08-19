package util

import (
	"unicode"
	"fmt"
)

type TokenType int32

const (
	LPAREN TokenType = iota
	RPAREN
	LAMBDA
	TERM
	EMPTY
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
	} else if t == EMPTY {
		return "EOF"
	} else if t == DOT {
		return "."
	} else {
		return "ERROR!"
	}
}

type Token struct {
	index uint64
	tokenType TokenType
	value string
}

func newToken(index uint64, tokenType TokenType, value string) Token {
	if tokenType != TERM {
		return Token{index, tokenType, ""}
	} else {
		return Token{index, tokenType, value}
	}
}


func Tokenize(input string) []Token {
	var result []Token
	current_term := ""
	char_position := 0

	if len(input) == 0 {
		return result
	}

	for _, ch := range input {
		char_position += 1
		next_token := Token{0, EMPTY, ""}
		if ch == 'λ' {
			next_token = newToken(uint64(char_position), LAMBDA, "")
		} else if  ch == '(' {
			next_token = newToken(uint64(char_position), LPAREN, "")
		} else if ch == '.' {
			next_token = newToken(uint64(char_position), DOT, "")
		} else if unicode.IsSpace(ch) {
			// Do nothing
		} else if ch == ')' {
			next_token = newToken(uint64(char_position), RPAREN, "")
		} else {
			current_term = current_term + string(ch)
			continue
		}

		if len(current_term) > 0 {
			result = append(result, newToken(uint64(char_position), TERM, current_term))
			current_term = ""
		}

		if next_token.tokenType != EMPTY {
			result = append(result, next_token)
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
			fmt.Print(item.value)
		} else if item.tokenType == EMPTY {
			fmt.Print(" ")
		} else if item.tokenType == DOT {
			fmt.Print(".")
		} else {
			fmt.Print("ERROR!!!!!")
		}
	}
	fmt.Print("\n")
}
