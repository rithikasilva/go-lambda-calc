package util

import (
	"fmt"
	"strconv"
)


func Interpret(e Expression) (Expression) {
	isNormalForm := false
	for !isNormalForm {
		reducer := newBetaReduction()
		reducedExpression := reducer.visit(e)
		isNormalForm = !reducer.reduced
		// Print the step
		fmt.Println(e.DisplayExpression())
		e = reducedExpression
	}
	return e
}






func visitFreeVariables(e Expression) map[string]bool {
	if e.getExpressionType() == VARIABLE {
		if varCast, ok := e.(Variable); ok {
			return createSet(varCast.term)
		} else {
			return createSet("error!")
		}
	} else if e.getExpressionType() == ABSTRACTION {
		if absCast, ok := e.(Abstraction); ok {
			left := createSet(absCast.term)
			right := visitFreeVariables(absCast.functionDefinition)
			return setDifference(left, right)
		} else {
			return createSet("error!")
		}		
	} else if e.getExpressionType() == APPLICATION {
		if appCast, ok := e.(Application); ok {
			left := visitFreeVariables(appCast.function)
			right := visitFreeVariables(appCast.argument)
			return setUnion(left, right)
		} else {
			return createSet("ERROR!")
		}
	} else {
		return createSet("error!")
	}
}



type AlphaConversion struct {
	toReplace string
	replacement Expression
}

func newAlphaConversion(tr string, r Expression) (*AlphaConversion) {
	return &AlphaConversion{tr, r}
}


func newName(unavailableNames map[string]bool, name string) string {
	num := 1
	for {
		currentName := name + strconv.Itoa(num)
		if _, ok := unavailableNames[currentName]; ok {
			num += 1
		} else {
			return currentName
		}
	}
}

func (a* AlphaConversion) visit(e Expression) (Expression) {
	if e.getExpressionType() == VARIABLE {
		if varCast, ok := e.(Variable); ok {
			if varCast.term == a.toReplace {
				return a.replacement
			} else {
				return newVariable(varCast.term)
			}
		} else {
			return newVariable("ERROR")
		}
	} else if e.getExpressionType() == ABSTRACTION {
		if absCast, ok := e.(Abstraction); ok {
			freeVars := visitFreeVariables(a.replacement)
			// Check to see if current thing is free or not
			if _, ok := freeVars[absCast.term]; ok {
				forUnvailable := visitFreeVariables(e)
				unavilableNames := setUnion(forUnvailable, createSet(absCast.term))
				newName := newName(unavilableNames, absCast.term)
				newFunctionDefinition := newVariable(newName)
				converter := newAlphaConversion(absCast.term, newFunctionDefinition)
				newBody := converter.visit(absCast.functionDefinition)
				ret := newAbstraction(newName, a.visit(newBody))
				return ret

			} else {
				return newAbstraction(absCast.term, a.visit(absCast.functionDefinition))
			}
		} else {
			return newVariable("ERROR")
		}
	} else if e.getExpressionType() == APPLICATION {
		if appCast, ok := e.(Application); ok {
			return newApplication(a.visit(appCast.function), a.visit(appCast.argument))
		} else {
			return newVariable("ERROR")
		}
	} else {
		return newVariable("ERROR")
	}
}


type BetaReduction struct {
	reduced bool
}

func newBetaReduction() (*BetaReduction) {
	return &BetaReduction{false}
}

func (b* BetaReduction) visit(e Expression) (Expression) {
	if e.getExpressionType() == VARIABLE {
		if varCast, ok := e.(Variable); ok {
			return newVariable(varCast.term)
		} else {
			return newVariable("ERROR")
		}
	} else if e.getExpressionType() == ABSTRACTION {
		if absCast, ok := e.(Abstraction); ok {
			return newAbstraction(absCast.term, absCast.functionDefinition)
		} else {
			return newVariable("ERROR")
		} 
	} else if e.getExpressionType() == APPLICATION {
		if appCast, ok := e.(Application); ok {
			if appCast.function.getExpressionType() == ABSTRACTION && !b.reduced {
				if absCast, ok := appCast.function.(Abstraction); ok {
					b.reduced = true
					converter := newAlphaConversion(absCast.term, appCast.argument)
					ret := converter.visit(absCast.functionDefinition)
					return ret
				} else {
					return newVariable("ERROR")
				}
			} else {
				return newApplication(b.visit(appCast.function), b.visit(appCast.argument))
			}
		} else {
			return newVariable("ERROR")
		}
	} else {
		return newVariable("ERROR")
	}
}
