- we need to do is repr our source code in other forms that are eaier to work with.![[Screenshot 2024-02-22 at 12.41.01 AM.png]]
- This transformation from source code to tokens is called as lexical analysis or lexing for short. Its done by a lexer.
- Tokens are itself are small,easily categorizable data structures that are then fed to the parser , which does the second transformation and turns the tokens into an AST.
- whitespaces characters dont show up as tokens. In our case , that's okay, because whitespace length is not significant in Monkey.
- TokenType type is defined as a string. Allows us to use many different values as TokenType , which in turn allows us to distinguish between different types of tokens. Using string also has the advantages of being easy to debug without a lot of boilerplate and helper functions: we can just print a string.
- we also defines the constants for the limited number of different token types in the Monkey language.
- Illegal signifies a token/character we dont know about and EOF stands for End of file which tells our parser later on that it can stop.

## The Lexer:
- Take source code as input and output the tokens that represent the source code. It will go through its input and output the next token it recognizes. It doesnt need to buffer or save tokens, since there will be only be one method called NextToken().
- Init the lexer with our source code and then repeatedly call NextToken() on it to go through the source code , token by token , character by character.
- doing a TDD based approach.

