package main

import (
	"bufio"
	"os"
	"fmt"
	"go-lambda-calc/util"
)


func tempTest() {
	content := []string{"((λx.x y) z)\n",
	"((λx.x y) z)\n",
	"(λx.x z)\n",
	"λx.x\n",
	"(λm.((m λt.λf.t) λx.λt.λf.f) λz.λs.z)\n",
	"(λx.λy.(x y) y)\n", "λx\n"}

	for _, item := range content {
		fmt.Println(item)
		result := util.Tokenize(item)
		p := util.NewParser(result)
		expression, err := p.Parse()
		if err != nil {
			fmt.Println("ERROR: ")
			fmt.Println(err)
		} else {
			fmt.Println("RESULT: ")
			fmt.Println(expression.DisplayExpression())
		}
		fmt.Print("\n\n")
	}
}


func lexAndParseMode() {
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
			if err != nil {
				fmt.Println("ERROR: ")
				fmt.Println(err)
			} else {
				fmt.Println("RESULT: ")
				fmt.Println(expression.DisplayExpression())
			}
		}
	}
}

func main() {
	lexAndParseMode()
}


/*
The rules I will follow:
Exp -> Var | App | Abs
Var -> TERM
App -> (Exp Exp)
Abs -> LAMBDA ID DOT Expression
*/
