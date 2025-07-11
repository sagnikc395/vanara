from .token import Token, TokenType

class Lexer:
    def __init__(self,input_):
        self.input = input_
        self.position = 0

    def next_token(self):
        if self.position >= len(self.input):
            return Token(TokenType.EOF,"")

        ch = self.input[self.position]
        self.position += 1

        token_map = {
            '=': TokenType.ASSIGN,
            '+': TokenType.PLUS,
            '(': TokenType.LPAREN,
            ')': TokenType.RPAREN,
            '{': TokenType.LBRACE,
            '}': TokenType.RBRACE,
            ',': TokenType.COMMA,
            ';': TokenType.SEMICOLON,
        }

        token_type = token_map.get(ch,TokenType.ILLEGAL)
        return Token(token_type,ch)
    
