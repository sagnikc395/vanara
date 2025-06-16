package ast

import "github.com/sagnikc395/vanara/pkg/token"

type ReturnStatment struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatment) statementNode() {}
func (rs *ReturnStatment) TokenLiteral() string {
	return rs.Token.Literal
}
