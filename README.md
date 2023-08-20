# go-lambda-calc
A fun little lambda calculus interpreter written in Go using only the standard library.

### Grammar Rules
All caps signifies terminals, others are non-terminals.
```
Expression -> Variable | Application | Abstraction
Variable -> TERM
Application -> LPAREN Expression Expression RPAREN
Abstraction -> LAMBDA ID DOT Expression

TERM: A unicode string.
LAMBDA: λ
DOT: .
LPAREN: (
RPAREN: )
```



### Approach
Execution uses normal order reduction. As the interpreter reduces, it prints each step on a separate line. Variable capture is also handled during α-conversion. Type `exit` to exit. Anything else is interpreted as an expression to be interpreted.


### TODO
- [ ] Allow setting custom terms that are interpreted as particular expressions during lexing.