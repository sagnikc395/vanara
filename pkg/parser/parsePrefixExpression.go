package parser

import "github.com/sagnikc395/vanara/pkg/ast"

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
	}

	p.NextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}
