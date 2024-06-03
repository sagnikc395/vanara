package ast

import "go/token"

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:token.LET,Literal: "let",
				},
			}
		},
	}
}