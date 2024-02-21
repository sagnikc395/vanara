package ast

import "monkey/token"

// ever node in AST has to implement this
// @return literal value of the token associated with
type Node interface {
	// TokenLiteral only used for debugging and testing
	TokenLiteral() string
}

// Some node(s) implement Statement
type Statement interface {
	Node
	statementNode()
}

// Some node(s) implement the Expression
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let statement would have 3 fields one for the identifier, one for the expression that produces the value and one for the token
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// identifier statement would have 2 fields the token and the value
type Identifier struct {
	Token token.Token
	Value string
}
