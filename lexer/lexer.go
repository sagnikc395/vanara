package lexer

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/sagnikc395/vanara/token"
)

type Lexer struct {
	//current character position
	position int
	//next character position
	readPosition int
	//current character
	ch rune
	//previous token
	prevToken token.Token
	//rune slice of our input string
	characters []rune
}

// constructor for lexer
func New(input string) *Lexer {
	l := &Lexer{characters: []rune(input)}
	l.readChar()
	return l
}

// getline returns the rough the line-number of our current position
func (l *Lexer) GetLine() int {
	line := 0
	chars := len(l.characters)
	i := 0

	for i < l.readPosition && i < chars {
		if l.characters[i] == rune('\n') {
			line++
		}
		i++
	}
	return line
}

// read one character forward
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.characters) {
		l.ch = 0
	} else {
		l.ch = l.characters[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	//skip whitespace characters
	l.skipWhitespace()

	//skip single line comment s
	if l.ch == rune('#') ||
		(l.ch == rune('/') && l.peekChar() == rune('/')) {
		l.skipComment()
		return (l.NextToken())
	}

	// multi-line comments
	if l.ch == rune('/') && l.peekChar() == rune('*') {
		l.skipMultiLineComment()
	}

	switch l.ch {
	case rune('&'):
		if l.peekChar() == rune('&') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AND, Literal: string(ch) + string(l.ch)}
		}
	case rune('|'):
		if l.peekChar() == rune('|') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.OR, Literal: string(ch) + string(l.ch)}
		}

	case rune('='):
		tok = newToken(token.ASSIGN, byte(l.ch))
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, byte(l.ch))
		}
	case rune(';'):
		tok = newToken(token.SEMICOLON, byte(l.ch))
	case rune('?'):
		tok = newToken(token.QUESTION, byte(l.ch))
	case rune('('):
		tok = newToken(token.LPAREN, byte(l.ch))
	case rune(')'):
		tok = newToken(token.RPAREN, byte(l.ch))
	case rune(','):
		tok = newToken(token.COMMA, byte(l.ch))
	case rune('.'):
		if l.peekChar() == rune('.') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.DOTDOT, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.PERIOD, byte(l.ch))
		}
	case rune('+'):
		if l.peekChar() == rune('+') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUS_PLUS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUS_EQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.PLUS, byte(l.ch))
		}

	case rune('%'):
		tok = newToken(token.MOD, byte(l.ch))
	case rune('{'):
		tok = newToken(token.LBRACE, byte(l.ch))
	case rune('}'):
		tok = newToken(token.RBRACE, byte(l.ch))
	case rune('-'):
		if l.peekChar() == rune('-') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUS_MINUS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUS_EQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.MINUS, byte(l.ch))
		}
	case rune('/'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.SLASH_EQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			// slash is mostly division, but could
			// be the start of a regular expression

			// We exclude:
			//   a[b] / c       -> RBRACKET
			//   ( a + b ) / c   -> RPAREN
			//   a / c           -> IDENT
			//   3.2 / c         -> FLOAT
			//   1 / c           -> IDENT
			//
			if l.prevToken.Type == token.RBRACKET ||
				l.prevToken.Type == token.RPAREN ||
				l.prevToken.Type == token.IDENT ||
				l.prevToken.Type == token.INT ||
				l.prevToken.Type == token.FLOAT {

				tok = newToken(token.SLASH, byte(l.ch))
			} else {
				str, err := l.readRegexp()
				if err == nil {
					tok.Type = token.REGEXP
					tok.Literal = str
				} else {
					tok.Type = token.REGEXP
					tok.Literal = str
				}
			}
		}
	case rune('*'):
		if l.peekChar() == rune('*') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.POW, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.ASTERISK_EQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASTERISK, byte(l.ch))
		}
	case rune('<'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LT_EQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LT, byte(l.ch))
		}
	case rune('>'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GT_EQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.GT, byte(l.ch))
		}
	case rune('~'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.CONTAINS, Literal: string(ch) + string(l.ch)}
		}

	case rune('!'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			if l.peekChar() == rune('~') {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.NOT_CONTAINS, Literal: string(ch) + string(l.ch)}

			} else {
				tok = newToken(token.BANG, byte(l.ch))
			}
		}
	case rune('"'):
		str, err := l.readString('"')
		if err == nil {
			tok.Literal = str
			tok.Type = token.STRING
		} else {
			tok.Literal = err.Error()
			tok.Type = token.ILLEGAL
		}
	case rune('`'):
		str, err := l.readString('`')
		if err == nil {
			tok.Literal = str
			tok.Type = token.BACKTICK
		} else {
			tok.Literal = err.Error()
			tok.Type = token.ILLEGAL
		}

	case rune('['):
		tok = newToken(token.LBRACKET, byte(l.ch))
	case rune(']'):
		tok = newToken(token.RBRACKET, byte(l.ch))
	case rune(':'):
		tok = newToken(token.COLON, byte(l.ch))
	case rune(0):
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isDigit(l.ch) {
			tok = l.readDecimal()
			l.prevToken = tok
			return tok

		}

		// Not printable?  That's a bug
		if !unicode.IsPrint(l.ch) {
			tok.Literal = string(l.ch)
			tok.Type = token.ILLEGAL

			// skip the characters
			l.readChar()
			return tok
		}

		tok.Literal = l.readIdentifier()

		// Did we fail to read a token?
		if len(tok.Literal) == 0 {
			// Then we've got an illegal
			tok.Type = token.ILLEGAL
			l.readChar()
			return tok
		}
		tok.Type = token.LookupIdentifier(tok.Literal)
		l.prevToken = tok

		return tok
	}

	l.readChar()
	l.prevToken = tok
	return tok
}

// return the new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier is designed to read an identifier
func (l *Lexer) readIdentifier() string {
	//functions valid with having dots in the name
	valid := map[string]bool{
		"directory.glob":     true,
		"math.abs":           true,
		"math.random":        true,
		"math.sqrt":          true,
		"os.environment":     true,
		"os.getenv":          true,
		"os.setenv":          true,
		"string.interpolate": true,
	}

	//types which will have valid methods
	types := []string{"string.", "array.", "integer.", "float.", "hash.", "object."}

	id := ""

	//save posn in case we need to jump backwards in our scanning
	position := l.position
	rposition := l.readPosition

	//build up our identifier, handling only valid characters

	for isIdentifier(l.ch) {
		id += string(l.ch)
		l.readChar()
	}

	//now we have to see if our identifier had a period inside it

	if strings.Contains(id, ".") {
		ok := valid[id]
		//check type-prefix, let the definition succeed
		if !ok {
			for _, i := range types {
				if strings.HasPrefix(id, i) {
					ok = true
				}
			}
		}

		//if not , then abort
		if !ok {
			//first truncate our identifier at the position before "."
			offset := strings.Index(id, ".")
			id = id[:offset]

			//move backwards
			l.position = position
			l.readPosition = rposition
			for offset > 0 {
				l.readChar()
				offset--
			}
		}
	}
	return id
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// skip comment -> until the end of the line
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != rune(0) {
		l.readChar()
	}
	l.skipWhitespace()
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// consume all the tokens until we had the close of a multi-line comment
func (l *Lexer) skipMultiLineComment() {
	found := false

	for !found {
		//break at end of our input
		if l.ch == rune(0) {
			found = true
		}

		//else, keep going until we find "*/"
		if l.ch == '*' && l.peekChar() == '/' {
			found = true
			//our current posn is "*", so skip forward to consume the "/"
			l.readChar()
		}
		l.readChar()
	}
	l.skipWhitespace()
}

// read number -> accept a digits, accept hex and binary
func (l *Lexer) readNumber() string {
	str := ""
	accept := "0123456789"

	//but if we have `0x` as a prefix we accept hexadecimal instead
	if l.ch == '0' && l.peekChar() == 'x' {
		accept = "0x123456789abcdefABCDEF"
	}

	//if we have `0b` as a prefix we accept binary digits only
	if l.ch == '0' && l.peekChar() == 'b' {
		accept = "b01"
	}

	for strings.Contains(accept, string(l.ch)) {
		str += string(l.ch)
		l.readChar()
	}

	return str
}

// read a decimal
func (l *Lexer) readDecimal() token.Token {
	//read a integer-number
	integer := l.readNumber()

	//either now
	// .[digit] -> int to float
	// .etc -> method call on raw number
	if l.ch == rune('.') && isDigit(l.peekChar()) {
		//float case
		l.readChar()
		frac := l.readNumber()
		return token.Token{Type: token.FLOAT, Literal: integer + "." + frac}
	}
	return token.Token{Type: token.INT, Literal: integer}
}

//read a string input ,delimited by the given character

func (l *Lexer) readString(delim rune) (string, error) {
	out := ""
	for {
		l.readChar()

		if l.ch == rune(0) {
			return "", fmt.Errorf("unterminated string")
		}

		if l.ch == delim {
			break
		}

		//handle \n,\r,\t,\"
		if l.ch == '\\' {
			//line ending with the \ + new line
			if l.peekChar() == '\n' {
				l.readChar()
				continue
			}

			l.readChar()

			if l.ch == rune(0) {
				return "", errors.New("unterminated string")
			}
			if l.ch == rune('n') {
				l.ch = '\n'
			}
			if l.ch == rune('r') {
				l.ch = '\r'
			}
			if l.ch == rune('t') {
				l.ch = '\t'
			}
			if l.ch == rune('"') {
				l.ch = '"'
			}
			if l.ch == rune('\\') {
				l.ch = '\\'
			}
		}
		out = out + string(l.ch)
	}
	return out, nil
}

// isDigit
func isDigit(ch rune) bool {
	return rune('0') <= ch && ch <= rune('9')
}

// peek the next character
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.characters) {
		return rune(0)
	} else {
		return l.characters[l.readPosition]
	}
}

// is white space
func isWhitespace(ch rune) bool {
	return ch == rune(' ') || ch == rune('\t') || ch == rune('\n') || ch == rune('\r')
}

// reading a regex function
func (l *Lexer) readRegexp() (string, error) {
	out := ""

	for {
		l.readChar()
		if l.ch == rune(0) {
			return "Unterminated regular expression", fmt.Errorf("Unterminated regular expression")
		}
		if l.ch == '/' {
			//consume termianting '/'
			l.readChar()

			//check flags
			flags := ""

			//two flags ; i -> Ignore-case , m -> Multiline
			for l.ch == rune('i') || l.ch == rune('m') {
				//save the char -> unless it is a repeat
				if !strings.Contains(flags, string(l.ch)) {
					tmp := strings.Split(flags, "")
					tmp = append(tmp, string(l.ch))
					flags = strings.Join(tmp, "")
				}
				//read the next
				l.readChar()
			}
			//convert the regexp to go
			if len(flags) > 0 {
				out = "(?" + flags + ")" + out
			}
			break
		}
		out = out + string(l.ch)
	}
	return out, nil
}

// determine ch is identifier or not
func isIdentifier(ch rune) bool {
	if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '.' || ch == '?' || ch == '$' || ch == '_' {
		return true
	}
	return false
}
