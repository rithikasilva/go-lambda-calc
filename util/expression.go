package util

import "fmt"

type ExpressionType uint16

const (
	VARIABLE ExpressionType = iota
	ABSTRACTION
	APPLICATION
)

type Expression interface {
	DisplayExpression() string
	getExpressionType() ExpressionType
}

type Variable struct {
	term           string
	expressionType ExpressionType
}

func newVariable(term string) Variable {
	return Variable{term, VARIABLE}
}

func (v Variable) getExpressionType() ExpressionType {
	return v.expressionType
}

func (v Variable) DisplayExpression() string {
	return v.term
}

type Abstraction struct {
	term               string
	functionDefinition Expression
	expressionType     ExpressionType
}

func newAbstraction(term string, rest Expression) Abstraction {
	return Abstraction{term, rest, ABSTRACTION}
}

func (a Abstraction) getExpressionType() ExpressionType {
	return a.expressionType
}

func (a Abstraction) DisplayExpression() string {
	return fmt.Sprintf("λ%s.%s", a.term, a.functionDefinition.DisplayExpression())
}

type Application struct {
	function       Expression
	argument       Expression
	expressionType ExpressionType
}

func newApplication(function Expression, argument Expression) Application {
	return Application{function, argument, APPLICATION}
}

func (a Application) getExpressionType() ExpressionType {
	return a.expressionType
}

func (a Application) DisplayExpression() string {
	function := a.function.DisplayExpression()
	argument := ""
	if a.argument.getExpressionType() == ABSTRACTION {
		argument = a.argument.DisplayExpression()
	} else if a.argument.getExpressionType() == APPLICATION {
		argument = fmt.Sprintf("(%s)", a.argument.DisplayExpression())
	} else {
		argument = a.argument.DisplayExpression()
	}
	return fmt.Sprintf("(%s %s)", function, argument)
}
