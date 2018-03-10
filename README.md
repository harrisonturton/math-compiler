# Math Compiler [![Travis](https://travis-ci.org/harrisonturton/math-compiler.svg?branch=master)](https://travis-ci.org/harrisonturton/math-compiler)

This project compiles arithmetic expressions into ARM assembly.

Supports:
- Addition
- Subtraction
- Multiplcation
- Division
- Exponents
- Nested Expressions

**Roadmap**

- [X] Create a non-ambiguous LL(k) grammar
- [x] Scanner
- [x] Recursive descent parser
- [x] IR code generation (ARM Assembly)
    - [x] Addition / Subtraction
    - [X] Exponents
    - [X] Integer Division
- [ ] Testing
    - [x] Scanner (61.9% coverage)
    - [x] Parser (72.6% coverage)
    - [ ] Codegen
- [ ] Error Messages




## Details

This is a learning project to implement the major parts of a compiler. It's overkill for simple math, but allows us to understand the full dev pipeline.

**Features:**

* Full lexical analyzer
	* based on [Rob Pike's presentation](https://www.youtube.com/watch?v=HxaD_trXwRE)
* Recursive Descent Parser
	* Parses concurrently with the lexer
* Assembly code generation


## Installation

```bash
$ git clone https://github.com/harrisonturton/math-compiler
```

This project does not require any special dependencies.



## Usage

These are shell scripts to examine the behaviour of the scanner/parser.

Each script can be passed a string, or a file.

#### Scanning

```
$ go run /src/scripts/scan.go -m "1+2"
TOK_NUMBER 1
TOK_ADD +
TOK_NUMBER 2
TOK_EOF
```

```
$ go run /src/scripts/scan.go filename.txt
TOK_NUMBER 1
TOK_ADD +
TOK_NUMBER 2
TOK_EOF
```

#### Parsing

```shell
$ go run /src/scripts/parse.go -m "1+2"
(+ 1 2)
```

```shell
$ go run /src/scripts/parse.go filename.txt
(+ 1 (* 3 4))
```

#### Compiling

```shell
$ go run /src/scripts/compile.go -m "1+(3-2)"
MOV r0, 1
PUSH {r0}
MOV r0, 3
PUSH {r0}
MOV r1, 2
PUSH {r1}
POP {r0, r1}
SUB r0, r1, r0
PUSH {r0}
POP {r0, r1}
ADD r0, r1
PUSH {r0}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
