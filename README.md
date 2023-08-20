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
Execution uses normal order reduction. As the interpreter reduces, it prints each step on a separate line. Variable capture is also handled during α-conversion.