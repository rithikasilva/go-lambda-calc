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


### Options
- Type expression to interpret to just interpret.
- Type `a` then `{X}={Y}` where `X` is a sequence of characters `A-z` as the key to subsitute and `Y` as the lambda expression to substitute in.
- Tyoe `q` to quit.


### Example expressions:
```
λx.x
(λx.x y)
((λx.x y) z)
(λx.λy.(x y) y)
(λa.((a λb.λc.b) λx.λb.λc.c) λd.λe.d)
```