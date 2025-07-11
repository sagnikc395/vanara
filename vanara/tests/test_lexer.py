from vanara.token import Token, TokenType
from vanara.lexer import Lexer  
import pytest


@pytest.mark.parametrize("expected_type, expected_literal", [
    (TokenType.ASSIGN, "="),
    (TokenType.PLUS, "+"),
    (TokenType.LPAREN, "("),
    (TokenType.RPAREN, ")"),
    (TokenType.LBRACE, "{"),
    (TokenType.RBRACE, "}"),
    (TokenType.COMMA, ","),
    (TokenType.SEMICOLON, ";"),
    (TokenType.EOF, ""),
])


def test_next_token(expected_type,expected_literal):
    input_text = "=+(){},;"
    lexer = Lexer(input_text)

    # cache tokens for all tests
    if not hasattr(test_next_token, "tokens"):
        test_next_token.tokens = [lexer.next_token() for _ in range(9)]

    index = test_next_token.call_count
    token = test_next_token.tokens[index]
    test_next_token.call_count += 1

    assert token.type == expected_type, f"token[{index}] type wrong: expected {expected_type}, got {token.type}"
    assert token.literal == expected_literal, f"token[{index}] literal wrong: expected {expected_literal}, got {token.literal}"

        
#init call count
test_next_token.call_count = 0 
