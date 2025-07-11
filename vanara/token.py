# our token and TokenType type

import enum


# defining tokentypes as enums    
class TokenType(enum.Enum):
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    # identifiers + literals
    IDENT = "IDENT"
    INT = "INT"

    # operators
    ASSIGN = "="
    PLUS = "+"

    #delimiters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    # keywords
    FUNCTION = "FUNCTION"
    LET = "LET"

class Token:
    type : TokenType
    literal : str    
