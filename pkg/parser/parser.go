package parser

import (
	"github.com/sagnikc395/vanara/pkg/ast"
	"github.com/sagnikc395/vanara/pkg/lexer"
	"github.com/sagnikc395/vanara/pkg/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
	errors    []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

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
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.NextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// error methods
func (p *Parser) Errors() []string {
	return p.errors
}
