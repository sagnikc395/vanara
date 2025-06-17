package ast

import "github.com/sagnikc395/vanara/pkg/token"

type ExpressionStatement struct {
	Token      token.Token // first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// stringer method
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
