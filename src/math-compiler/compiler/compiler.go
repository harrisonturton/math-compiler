package compiler

import (
	"../parser"
	"../token"
	"fmt"
)

// Generate target assembly for AST specified in /parser/ast
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

// Generate target code for a Number
func compileNumber(num parser.Number) string {
	return fmt.Sprintf("MOV r0, %s\nPUSH {r0}", num.Token.Value)
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
		result += "\nPUSH {r0}"
	}
	return result
}

// Generate target code for a Binary Op
func compileBinaryOp(binaryOp parser.BinaryOp) string {
	result := ""
	result += compileBinaryLeaf(binaryOp.Left, true)
	result += compileBinaryLeaf(binaryOp.Right, false)
	result += "\nPOP {r0, r1}"
	switch binaryOp.Op.Token {
	case token.TOK_ADD:
		result += "\nADD r0, r1"
	case token.TOK_SUB:
		result += "\nSUB r0, r1"
	}
	result += "\nPUSH {r0}"
	return result
}

func compileBinaryLeaf(leaf parser.Expr, isLeft bool) string {
	register := "r0"
	if !isLeft {
		register = "r1"
	}
	switch leaf.(type) {
	case parser.Number:
		num := leaf.(parser.Number)
		return fmt.Sprintf("\nMOV %s, %s", register, num.Token.Value)
	case parser.UnaryOp:
		unaryOp := leaf.(parser.UnaryOp)
		return compileUnaryOp(unaryOp)
	case parser.BinaryOp:
		binaryOp := leaf.(parser.BinaryOp)
		return compileBinaryOp(binaryOp)
	}
	return ""
}
