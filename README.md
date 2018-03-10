# Math Compiler [![Travis](https://travis-ci.org/harrisonturton/math-compiler.svg?branch=master)](https://travis-ci.org/harrisonturton/math-compiler)

This project compiles any arithmetic expression into ARM assembly.

**Not production-ready.** This is for me to learn the compiler pipeline.

- Full lexical analyser
  - Based on  [Rob Pike's presentation](https://www.youtube.com/watch?v=HxaD_trXwRE)
- Recursive Descent Parser
  - Lexes & Parses concurrently
- Assembly Code generation
- Full arithmetic support (excluding floating-point numbers)
  - Addition / Subtraction
  - Multiplication / Integer Division
  - Exponents
  - Nested expressions (Parenthesized)

## Installation

```bash
$ git clone https://github.com/harrisonturton/math-compiler
```

This project does not require any special dependencies.



## Usage

This project does not have a dedicated binary. Instead, it provides scripts to scan, parse, or compile input.

Long output has been trunctated using `…`.

#### Compiling

```bash
$ go run math-compiler/src/scripts/compile.go -m "1+(2^4)"
main:
  MOV r0, 1
...
branchToLR:
  bx lr
```

#### Parsing

```bash
$ go run math-compiler/src/scripts/parse.go -m "1+(2^4)"
(+ 1 (^ 2 3))
```

#### Compiling

```bash
$ go run math-compiler/src/scripts/scan.go -m "1+(2^4)"
TOK_NUMBER 1
TOK_ADD +
...
TOK_RPAREN )
TOK_EOF
```



## Roadmap

- [x] Unambiguous LL(k) grammar for arithmetic
- [x] Scanner (Lexer)
- [x] Recursive Descent Parser
- [x] IR code generation (ARM Assembly)
    - [x] Addition / Subtraction
    - [x] Multiplication / Division
    - [x] Exponents
- [ ] Full Testing
    - [x] Scanner (61.9% coverage)
    - [x] Parser (72.6% coverage)
    - [ ] Code generation
- [ ] Error Messages



## License

[MIT](https://choosealicense.com/licenses/mit/)

