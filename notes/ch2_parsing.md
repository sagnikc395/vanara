- Parser is as common as compiler, interpreter and programming language. 
- Everyone knows that parseres exist.
- But what is parser exactly?
  - parser is a software component that takes input data and build a data structure - often some kind of aprse tree , AST or other hierarchical structure - giving a structural repr of the input ,checking for correct syntax in the process. 
  - Parser often preceded by a seperate lexical analyzer , which creates tokens from the tokens of input characters.
  
- The output of parser , the AST is prettry abstract: there are no parentheses, no semicolons and no braces. But it does repr the source code pretty accurately.

- Parsers take source code as input (either text or tokens) and produce a data structure which represents this source code. While building up the data structure, they uanvoidably analyze the input, checking it confirms to the expected structure. Due to this it is also called as syntactic analysis.
  
### Parser Generators and why not for Monkey: (may change views in future :D)

- Parser Generators are tools that when fed with a formal description of a language , produce parsers as their output. This output is code that can then be compiled/ interpreter and itself fed with source code as input to produce a syntax tree.
- Eg: yacc, ANTLR etc.
- Majority of parser generators use a CFG(Context Free Grammar) as thier input. A CFG is a set of rules that describe how to form correct sentences in a language. The most common notation formats of CFGs are the BNF or the Extended BNF(Backus Naur Form).
- Parsing is exceptionally well suited to being automatically generated. Parsing is one of the most well-understood brnacehs of computer science and really smart people have already invested a lot of time into the problems of parsing.

### Parsing Techniques:
- 2 ways of parsing :
  - Top Down Parsing
  - Bottom Up Parsing
- Other methods:
  - RDP -> Recurive Descent Parsing 
  - Early Parsing 
  - Predictive Parsing 
  
- Here using RDP for now. Its a top-down operator precedence parser, sometimes called Pratt parser.
- Diff between top-down and bottom-up parsers:
  - Former starts with cosntructing root node of the AST and then descends while the latter does it the other way around.
  - A RDP parser, which works from the top-down , is often recommended for newcomers to paersing , since it closely mirrors the way we think about ASTs and their construction.

- Tradeoffs using RDP:
  - wont be the fastest :p
  - no formal proof for correctness and its error recovery process and detection of erroneous syntax wont be bullet proof.

### Parsing let statements:
- let statemenets :
  - bind a value to the given name.
  - our job here is to parse let statement correctly.
  - for now going to skip parsing the expressions that produce the value of a given variable binding and come abck to this later - as soon as we know how to parse expressions on their own.
- **what is parsing correctly ?**
  - The parser will produce an AST that accurately represent the information contained in original let statement.
- Programs in Monkey are a series of statement.
  - let \<identifier> = \<expression>
  - let statement in monkey consists of 2 changing parts: an identifier and an expression.
  - Diff b/w statements and expressions ?
    - Expressions produce values , statements dont.
    - let x = 5 -> no value, just an binding.
    - 5 produces a value -> 5 
    - return 5; -> no value, a statement 
    - add(5,5) -> value 
  - what is an expression or a statement , what produces value and what doesn't , depends on the programming language.