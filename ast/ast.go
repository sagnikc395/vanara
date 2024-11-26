package ast

import (
	"github.com/sagnikc395/vanara/token"
)

// implement by every node
type Node interface {
	//used for debugging and testing
	TokenLiteral() string
}

// some are statements
type Statement interface {
	Node
	statementNode()
}

// while others would be expressions
type Expression interface {
	Node
	expressionNode()
}

// root node of every AST our parser produces.
// series of statements and contained in the program, just a slice of AST nodes that implement the Statement interface
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

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

