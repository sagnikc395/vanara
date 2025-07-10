// define the basic lexer

use crate::token::{Token, TokenType};

pub struct Lexer {
    input: String,
    position: usize,      // current posn in input (points to current char)
    read_position: usize, // current reading posn in input ( after current char)
    ch: Option<char>,     //current char under examination
}

impl Lexer {
    pub fn new(input: String) -> Self {
        let mut l = Lexer {
            input,
            position: 0,
            read_position: 0,
            ch: None,
        };
        l.read_char();
        l
    }

    fn read_char(&mut self) {
        if self.read_position >= self.input.len() {
            self.ch = None;
        } else {
            self.ch = self.input.chars().nth(self.read_position);
        }
        self.position = self.read_position;
        self.read_position += 1;
    }

    fn skip_whitespace(&mut self) {
        while let Some(c) = self.ch {
            if c.is_whitespace() {
                self.read_char();
            } else {
                break;
            }
        }
    }
    pub fn next_token(&mut self) -> Token {
        //get the next token and return it
        self.skip_whitespace();

        let token = match self.ch {
            Some('=') => Token {
                token_type: TokenType::Assign,
                literal: "=".to_string(),
            },
            Some('+') => Token {
                token_type: TokenType::Plus,
                literal: "+".to_string(),
            },
            Some('(') => Token {
                token_type: TokenType::LParen,
                literal: "(".to_string(),
            },
            Some(')') => Token {
                token_type: TokenType::RParen,
                literal: ")".to_string(),
            },
            Some('{') => Token {
                token_type: TokenType::LBrace,
                literal: "{".to_string(),
            },
            Some('}') => Token {
                token_type: TokenType::RBrace,
                literal: "}".to_string(),
            },
            Some(',') => Token {
                token_type: TokenType::Comma,
                literal: ",".to_string(),
            },
            Some(';') => Token {
                token_type: TokenType::Semicolon,
                literal: ";".to_string(),
            },
            None => Token {
                token_type: TokenType::Eof,
                literal: "".to_string(),
            },
            _ => Token {
                token_type: TokenType::Illegal,
                literal: self.ch.unwrap().to_string(),
            },
        };
        self.read_char();
        token
    }
}
