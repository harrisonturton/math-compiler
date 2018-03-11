# Math Compiler [![Travis](https://travis-ci.org/harrisonturton/math-compiler.svg?branch=master)](https://travis-ci.org/harrisonturton/math-compiler)

This project compiles any arithmetic expression into ARM assembly.

Test this interactively on the [project website](https://harrisonturton.github.io/math-compiler/)

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

Go to https://harrisonturton.github.io/math-compiler/ to test this project interactively.

Alternatively, this project can be compiled and used from the commandline: (This requires you to [build it from source](#build-from-source))

#### Compiling

```bash
$ ./main "1+(2^4)"
main:
  MOV r0, 1
...
branchToLR:
  bx lr
```

## Build from Source

To use this project from the commandline, you'll need to build it.

Requirements:
* Working Go installation

Download the project as specified in [Installation](#installation), and place it in your `$GOROOT`.

In the project root, run:

```
$ cd build && go build ../src/main
```

A `main` binary should be created in the `/build` subdirectory. You can run this as so:

```
./main "1+2*3"
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
    - [X] Syntax Errors
    - [X] Basic parsing errors
    - [ ] Errors for unbalanced parenthesis



## License

[MIT](https://choosealicense.com/licenses/mit/)

