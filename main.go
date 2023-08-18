package main

import (
	"fmt"
)

type Expression interface {
	displayExpression() string
	getIdentity() string
}

type Variable struct {
	term string
	indentity string
}

func newVariable(term string) Variable {
	return Variable{term, "Variable"}
}

type Abstraction struct {
	term string
	functionDefinition Expression
	indentity string
}

func newAbstraction(term string, rest Expression) Abstraction {
	return Abstraction{term, rest, "Abstraction"}
}

type Application struct {
	function Expression
	argument Expression
	indentity string
}

func newApplication(function Expression, argument Expression) Application {
	return Application{function, argument, "Application"}
}


func (v Variable) getIdentity() string {
	return v.indentity
}

func (a Abstraction) getIdentity() string {
	return a.indentity
}

func (a Application) getIdentity() string {
	return a.indentity
}


func (v Variable) displayExpression() string {
	return v.term
}

func (a Abstraction) displayExpression() string {
	return "λ" + a.term + "." + a.functionDefinition.displayExpression() 
}

func (a Application) displayExpression() string {
	function := ""
	if a.function.getIdentity() == "Abstraction" {
		function = "(" + a.function.displayExpression() + ") "
	} else {
		function = a.function.displayExpression() + " "
	}
	argument := ""
	if a.argument.getIdentity() == "Abstraction" {
		argument = "(" + a.argument.displayExpression() + ")"
	} else if a.argument.getIdentity() == "Application" {
		argument = "(" + a.argument.displayExpression() + ")"
	} else {
		argument = a.argument.displayExpression()
	}
	return function + " " + argument
}




func main() {
	thing := newApplication(
		newAbstraction("x", newApplication(
			newVariable("x"), newVariable("y"))),
			 newVariable("z"))
	fmt.Println(thing.displayExpression())
}