package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// pre-defined types
const (
	AND             = "&&"
	ASSIGN          = "="
	ASTERISK        = "*"
	ASTERISK_EQUALS = "*="
	BACKTICK        = "`"
	BANG            = "!"
	CASE            = "case"
	COLON           = ":"
	COMMA           = ","
	CONST           = "CONST"
	CONTAINS        = "~="
	DEFAULT         = "DEFAULT"
	DEFINE_FUNCTION = "DEFINE_FUNCTION"
	DOTDOT          = ".."
	ELSE            = "ELSE"
	EOF             = "EOF"
	EQ              = "=="
	FALSE           = "FALSE"
	FLOAT           = "FLOAT"
	FOR             = "FOR"
	FOREACH         = "FOREACH"
	FUNCTION        = "FUNCTION"
	GT              = ">"
	GT_EQUALS       = ">="
	IDENT           = "IDENT"
	IF              = "IF"
	ILLEGAL         = "ILLEGAL"
	IN              = "IN"
	INT             = "INT"
	LBRACE          = "{"
	LBRACKET        = "["
	LET             = "LET"
	LPAREN          = "("
	LT              = "<"
	LT_EQUALS       = "<="
	MINUS           = "-"
	MINUS_EQUALS    = "-="
	MINUS_MINUS     = "--"
	MOD             = "%"
	NOT_CONTAINS    = "!~"
	NOT_EQ          = "!="
	NULL            = "null"
	OR              = "||"
	PERIOD          = "."
	PLUS            = "+"
	PLUS_EQUALS     = "+="
	PLUS_PLUS       = "++"
	POW             = "**"
	QUESTION        = "?"
	RBRACE          = "}"
	RBRACKET        = "]"
	REGEXP          = "REGEXP"
	RETURN          = "RETURN"
	RPAREN          = ")"
	SEMICOLON       = ";"
	SLASH           = "/"
	SLASH_EQUALS    = "/="
	STRING          = "STRING"
	SWITCH          = "switch"
	TRUE            = "TRUE"
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"case":     CASE,
	"const":    CONST,
	"default":  DEFAULT,
	"for":      FOR,
	"foreach":  FOREACH,
	"function": DEFINE_FUNCTION,
	"in":       IN,
	"null":     NULL,
	"switch":   SWITCH,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
