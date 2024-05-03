package parser

import (
	"github.com/sagnikc395/monkey/ast"
	"github.com/sagnikc395/monkey/lexer"
	"github.com/sagnikc395/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//reading the 2 tokens, so currToken and peekToken aer both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}


