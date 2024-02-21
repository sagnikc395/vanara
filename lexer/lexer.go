package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to the current character)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// return a new Lexer instance
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
