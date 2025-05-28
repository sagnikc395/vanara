package parser

import (
	"github.com/sagnikc395/vanara/ast"
	"github.com/sagnikc395/vanara/lexer"
	"github.com/sagnikc395/vanara/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//read the two tokens, so currTokens and peekTokens are both set
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
