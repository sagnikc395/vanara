#[derive(Debug, PartialEq, Clone)]
pub enum TokenType {
    Assign,
    Plus,
    LParen,
    RParen,
    LBrace,
    RBrace,
    Comma,
    Semicolon,
    Eof,
    //for unrecognized characters
    Illegal,
    Ident,
    Int,
}

#[derive(Debug, PartialEq, Clone)]
pub struct Token {
    pub token_type: TokenType,
    pub literal: String,
}
