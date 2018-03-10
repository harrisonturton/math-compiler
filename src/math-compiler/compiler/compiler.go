package compiler

import (
	"../parser"
	"../token"
	"fmt"
)

const exponent = `exponent:
  cmp r1, 0
  beq branchToLR
  mul r0, r2
  sub r1, 1
  b exponent
branchToLR:
  bx lr`

  func Compile(ast parser.Expr) string {
	  result := ""
	  result += compileRoot(ast)
	  result += "\nfinish:"
	  result += "\n  nop"
	  result += "\n  b finish"
	  result += "\n" + exponent
	  return result
  }

func compileRoot(ast parser.Expr) string {
	result := ""
	switch ast.(type) {
	case parser.Number:
		num := ast.(parser.Number)
		result += compileNumber(num)
	case parser.UnaryOp:
		unaryOp := ast.(parser.UnaryOp)
		result += compileUnaryOp(unaryOp)
	case parser.BinaryOp:
		binaryOp := ast.(parser.BinaryOp)
		result += compileBinaryOp(binaryOp)
	}
	return result
}

func compileNumber(num parser.Number) string {
	return fmt.Sprintf("\nMOV r0, %s\nPUSH {r0}", num.Token.Value)	
}

func compileUnaryOp(unaryOp parser.UnaryOp) string {
	result := ""
	switch unaryOp.Right.(type) {
	case parser.Number:
		num := unaryOp.Right.(parser.Number)
		result += compileNumber(num)
	case parser.UnaryOp:
		innerUnaryOp := unaryOp.Right.(parser.UnaryOp)
		result += compileUnaryOp(innerUnaryOp)
	case parser.BinaryOp:
		binaryOp := unaryOp.Right.(parser.BinaryOp)
		result += compileBinaryOp(binaryOp)
	}
	result += "\nPOP {r0}"
	switch unaryOp.Op.Token {
	case token.TOK_ADD:
		result += "\nPUSH {r0}"
	case token.TOK_SUB:
		result += "\nSUB r0, 0, r0"
	}
	return result
}

func compileBinaryOp(binaryOp parser.BinaryOp) string {
	result := ""
	result += compileRoot(binaryOp.Left)
	result += compileRoot(binaryOp.Right)
	switch binaryOp.Op.Token {
	case token.TOK_ADD:
		result += "\nPOP {r0, r1}"
		result += "\nADD r0, r1"
	case token.TOK_SUB:
		result += "\nPOP {r0, r1}"
		result += "\nSUB r0, r1, r0"
	case token.TOK_MUL:
		result += "\nPOP {r0, r1}"
		result += "\nMUL r0, r1"
	case token.TOK_DIV:
		result += "\nPOP {r0, r1}"
		result += "\nSDIV r0, r1, r0"
	case token.TOK_ORD:
		result += "\nPOP {r1, r2}"
		result += "\nMOV r0, 1"
		result += "\nbl exponent"
	}
	result += "\nPUSH {r0}"
	return result
}

