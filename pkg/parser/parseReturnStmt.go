package parser

import (
	"github.com/sagnikc395/vanara/pkg/token"

	"github.com/sagnikc395/vanara/pkg/ast"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatment {
	stmt := &ast.ReturnStatment{Token: p.currToken}

	p.NextToken()

	//TODO: currently skipping the expressions until we encounter a variable
	for !p.currTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}
