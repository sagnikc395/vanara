package lexer

type Lexer struct {
	input        string
	position     int  // curr posn in input
	readPosition int  // current reading posn in input
	ch           byte // current char under check
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
