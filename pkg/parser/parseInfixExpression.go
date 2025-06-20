package parser

import "github.com/sagnikc395/vanara/pkg/ast"

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
		Left:     left,
	}

	precedence := p.currPrecedence()
	p.NextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}
