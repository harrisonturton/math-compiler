# Math Compiler [![Travis](https://travis-ci.org/harrisonturton/math-compiler.svg?branch=master)](https://travis-ci.org/harrisonturton/math-compiler)

This project compiles simple arithmetic into ARM v7 assembly.

This is a learning project to understanding language development. It has a fully-featured lexical analyzer & a recursive descent parser.

It can currently scan & parse arithmetic expressions, but IR generation (Assembly) is in progress.

**Progress**

- [X] Create a non-ambiguous LL(k) grammar
- [X] Scanner
- [X] Tests for scanner
- [X] Recursive descent parser
- [X] Tests for parser
- [ ] IR code generation (ARM Assembly)
- [ ] Machine code generation (ARM v7)
- [ ] Error Messages



## Installation

```bash
$ git clone https://github.com/harrisonturton/math-compiler
```

This project does not require any special dependencies.



## Usage

Only scanning (lexing) and parsing are supported - compilation is still in development.

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

```bash
$ go run /src/scripts/parse.go -m "1+2"
(+ 1 2)
```

```
$ go run /src/scripts/parse.go filename.txt
(+ 1 (* 3 4))
```



## License

[MIT](https://choosealicense.com/licenses/mit/)
