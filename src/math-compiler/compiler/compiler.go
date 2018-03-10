package compiler

import (
	"../parser"
	"../token"
	"fmt"
)

func Compile(ast parser.Expr) string {
	switch ast.(type) {
	case parser.Number:
		num := ast.(parser.Number)
		return compileNumber(num)
	case parser.UnaryOp:
		unaryOp := ast.(parser.UnaryOp)
		return compileUnaryOp(unaryOp)
	case parser.BinaryOp:
		binaryOp := ast.(parser.BinaryOp)
		return compileBinaryOp(binaryOp)
	}
	return ""
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
	result += Compile(binaryOp.Left)
	result += Compile(binaryOp.Right)
	result += "\nPOP {r0, r1}"
	switch binaryOp.Op.Token {
	case token.TOK_ADD:
		result += "\nADD r0, r1"
	case token.TOK_SUB:
		result += "\nSUB r0, r1, r0"
	case token.TOK_MUL:
		result += "\nMUL r0, r1"
	case token.TOK_DIV:
		result += "\nSDIV r0, r1, r0"
	}
	result += "\nPUSH {r0}"
	return result
}

