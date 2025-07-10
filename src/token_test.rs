#[cfg(test)]
mod tests {
    use super::super::token::{Token, TokenType};
    use super::Lexer;

    #[test]
    fn text_next_token() {
        let input = "=+(){},;".to_string();

        let tests = vec![
            (TokenType::Assign, "=".to_string()),
            (TokenType::Plus, "+".to_string()),
            (TokenType::LParen, "(".to_string()),
            (TokenType::RParen, ")".to_string()),
            (TokenType::LBrace, "{".to_string()),
            (TokenType::RBrace, "}".to_string()),
            (TokenType::Comma, ",".to_string()),
            (TokenType::Semicolon, ";".to_string()),
            (TokenType::Eof, "".to_string()),
        ];

        let mut l = Lexer::new(input);

        for (expected_type, expected_literal) in tests {
            let tok = l.next_token();

            assert_eq!(
                tok.token_type, expected_type,
                "Wrong token type. Expected {:?}, got {:?}",
                expected_type, tok.token_type
            );
            assert_eq!(
                tok.literal, expected_literal,
                "Wrong literal. Expected \"{}\", got \"{}\"",
                expected_literal, tok.literal
            );
        }
    }
}
