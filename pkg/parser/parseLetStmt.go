package parser

import (
	"github.com/sagnikc395/vanara/pkg/ast"
	"github.com/sagnikc395/vanara/pkg/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	//TODO: skipping the expression until we encounter a semicolon
	for !p.currTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}
