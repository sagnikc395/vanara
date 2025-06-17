package ast

import (
	"bytes"

	"github.com/sagnikc395/vanara/pkg/token"
)

type ReturnStatment struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatment) statementNode() {}
func (rs *ReturnStatment) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatment) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
