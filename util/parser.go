package util

import (
	"errors"
)

type Parser struct {
	data []Token
	next Token
}

func NewParser(t []Token) (*Parser) {
	return &Parser{t[1:], t[0]}
}

func (p* Parser) parseAdvance() {
	if len(p.data) == 0 {
		p.data = nil
		p.next = newToken(0, EMPTY, "EOF")
	} else {
		p.next = p.data[0]
		p.data = p.data[1:]
	}
}

func (p* Parser) parseConsume(t TokenType) (error) {
	if p.next.tokenType == t {
		p.parseAdvance()
		return nil
	} else {
		return errors.New(makeParseError(tokenTypeToString(t), tokenTypeToString(p.next.tokenType)))
	}
}

func (p* Parser) parseExpression() (Expression, error) {
	if p.next.tokenType == LPAREN {
		return p.parseApplication()
	} else if p.next.tokenType == LAMBDA {
		return p.parseAbstraction()
	} else if p.next.tokenType == TERM {
		return p.parseVariable()
	} else {
		return newVariable("invalid"), errors.New(makeParseError("(, λ, or a TERM", tokenTypeToString(p.next.tokenType)))
	}

}

func (p* Parser) parseVariable() (Expression, error) {
	if p.next.tokenType == TERM {
		name := p.next.value
		p.parseAdvance()
		return newVariable(name), nil
	} else {
		return newVariable("invalid"), errors.New(makeParseError("a TERM", tokenTypeToString(p.next.tokenType)))
	}
}

func (p* Parser) parseApplication() (Expression, error) {
	if p.next.tokenType == LPAREN {
		p.parseAdvance() // Consume the lambda
		leftExpression, err := p.parseExpression()
		if err != nil {
			return newVariable("invalid"), err
		}
		rightExpression, err := p.parseExpression()
		if err != nil {
			return newVariable("invalid"), err
		}
		p.parseConsume(RPAREN)
		return newApplication(leftExpression, rightExpression), nil
	} else {
		return newVariable("invalid"), errors.New(makeParseError("(", tokenTypeToString(p.next.tokenType)))
	}
}


func (p* Parser) parseAbstraction() (Expression, error) {
	if p.next.tokenType == LAMBDA {
		p.parseAdvance()
		variable, err := p.parseVariable()
		if err != nil {
			return newVariable("invalid"), err
		}
		if varCast, ok := variable.(Variable); ok {
			err := p.parseConsume(DOT)
			if err != nil {
				return newVariable("invalid"), err
			}
			expression, err :=  p.parseExpression()
			if err != nil {
				return newVariable("invalid"), err
			}
			return newAbstraction(Variable(varCast).term, expression), nil
		} else {
			return newVariable("invalid"), errors.New("internal error")
		}
	} else {
		return newVariable("invalid"), errors.New(makeParseError("λ", tokenTypeToString(p.next.tokenType)))
	}
}

func (p* Parser) Parse() (Expression, error) {
	return p.parseExpression()
}

func makeParseError(expected string, found string) string {
	return "Expected: " + expected + ", Found: " + found
}

