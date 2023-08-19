package util


type Expression interface {
	DisplayExpression() string
	getIdentity() string
}

type Variable struct {
	term string
	indentity string
}
func newVariable(term string) Variable {
	return Variable{term, "Variable"}
}
func (v Variable) getIdentity() string {
	return v.indentity
}
func (v Variable) DisplayExpression() string {
	return v.term
}


type Abstraction struct {
	term string
	functionDefinition Expression
	indentity string
}

func newAbstraction(term string, rest Expression) Abstraction {
	return Abstraction{term, rest, "Abstraction"}
}
func (a Abstraction) getIdentity() string {
	return a.indentity
}
func (a Abstraction) DisplayExpression() string {
	return "λ" + a.term + "." + a.functionDefinition.DisplayExpression() 
}

type Application struct {
	function Expression
	argument Expression
	indentity string
}

func newApplication(function Expression, argument Expression) Application {
	return Application{function, argument, "Application"}
}

func (a Application) getIdentity() string {
	return a.indentity
}

func (a Application) DisplayExpression() string {
	function := ""
	if a.function.getIdentity() == "Abstraction" {
		function = a.function.DisplayExpression() + " "
	} else {
		function = a.function.DisplayExpression() + " "
	}
	argument := ""
	if a.argument.getIdentity() == "Abstraction" {
		argument = a.argument.DisplayExpression()
	} else if a.argument.getIdentity() == "Application" {
		argument = " (" + a.argument.DisplayExpression() + ")"
	} else {
		argument = a.argument.DisplayExpression()
	}
	return "(" + function + argument + ")"
}