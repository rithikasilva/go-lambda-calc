package main

import (
	"bufio"
	"errors"
	"fmt"
	"go-lambda-calc/util"
	"os"
	"strings"
)

func isValidSubstitution(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func checkNewSubstitution(text string) (string, string, error) {
	splitted := strings.Split(text, "=")
	if len(splitted) != 2 {
		return "", "", errors.New("not text and corresponding substitution")
	}
	shorthand := splitted[0]
	if !isValidSubstitution(shorthand) {
		return "", "", errors.New("substitution text does not contains characters from A-z")
	}
	substitution := splitted[1]
	result := util.Tokenize(substitution, map[string]string{})
	parser := util.NewParser(result)
	_, err := parser.Parse()
	if err != nil {
		return "", "", err
	} else {
		return shorthand, substitution, nil
	}
}

func run() {
	reader := bufio.NewReader(os.Stdin)
	substitution := map[string]string{}
	for {
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		if len(text) == 1 {
			fmt.Println("ERROR: ")
			fmt.Println("invalid length for input")
		} else if text == "q\n" {
			break
		} else if text == "a\n" {
			text, _ := reader.ReadString('\n')
			key, value, err := checkNewSubstitution(text)
			if err != nil {
				fmt.Println(err)
			} else {
				substitution[key] = value
			}
			continue
		} else {
			result := util.Tokenize(text, substitution)
			parser := util.NewParser(result)
			expression, err := parser.Parse()
			fmt.Println("Parsed:")
			fmt.Println(expression.DisplayExpression())
			fmt.Println("---------")
			fmt.Println("REDUCTION:")
			reduced := util.Interpret(expression)
			fmt.Println("---------")
			if err != nil {
				fmt.Println("ERROR: ")
				fmt.Println(err)
			} else {
				fmt.Println("RESULT REDUCED: ")
				fmt.Println(reduced.DisplayExpression())
			}
		}
	}
}

func main() {
	run()
}
