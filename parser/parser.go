//construct ast with our parser

package parser

import (
	"github.com/sagnikc395/vanara/ast"
	"github.com/sagnikc395/vanara/lexer"
	"github.com/sagnikc395/vanara/token"
)

// l is a pointer to an instsance of the lexer , on which we repeatedly call NextToken() to get the next token in the input.
// curToken and peekToken like pointers, they point towards the current and the next token.
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//read the tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// entry part of parser and it will construct the root node of the AST.
// it then builds the child nodes, the statements, by calling other functions that know
// which AST node to construct based on the current token.
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
