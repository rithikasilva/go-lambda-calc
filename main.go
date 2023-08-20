package main

import (
	"bufio"
	"os"
	"fmt"
	"go-lambda-calc/util"
)


func tempTest() {
	content := []string{
	"((λx.x y) z)\n",
	"(λx.x z)\n",
	"λx.x\n",
	"(λm.((m λt.λf.t) λx.λt.λf.f) λz.λs.z)\n",
	"(λx.λy.(x y) y)\n", "λx\n",
	"exit\n"}

	for _, text := range content{
		fmt.Print("\n\n\n")
		if text == "exit\n" {
			break
		} else if len(text) == 1 {
			fmt.Println("ERROR: ")
			fmt.Println("invalid length for input")
			continue
		} else {
			result := util.Tokenize(text)
			parser := util.NewParser(result)
			expression, err := parser.Parse()
			fmt.Println("Parsed:")
			fmt.Println(expression.DisplayExpression())
			fmt.Println("---------")
			fmt.Println("REDUCTION:")
			reduced :=  util.Interpret(expression)
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


func lexParseInterpretMode() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		if text == "exit\n" {
			break
		} else if len(text) == 1 {
			fmt.Println("ERROR: ")
			fmt.Println("invalid length for input")
			continue
		} else {
			result := util.Tokenize(text)
			parser := util.NewParser(result)
			expression, err := parser.Parse()
			fmt.Println("Parsed:")
			fmt.Println(expression.DisplayExpression())
			fmt.Println("---------")
			fmt.Println("REDUCTION:")
			reduced :=  util.Interpret(expression)
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
	tempTest()
}

