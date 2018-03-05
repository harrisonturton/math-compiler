package compiler

import (
	"../parser"
	"../token"
	"fmt"
)

/*
	5/3
	MOV r0, 0
	MOV r1, 5
	MOV r2, 3
	cpmgt r1, r1
	ADD r0, 1
	SUB r1, r1
	b 15
*/

func Divide(value int, div int) int {
	result := 0
	for value >= div {
		result += 1
		value -= div
	}
	return result
}

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

// Handle logic for compiling each binary leaf.
// Need seperate logic since each leaf result
// needs to be stored in a different register.
func compileBinaryLeaf(leaf parser.Expr, isLeft bool) string {
	register := "r0"
	if !isLeft {
		register = "r1"
	}
	switch leaf.(type) {
	case parser.Number:
		num := leaf.(parser.Number)
		return fmt.Sprintf("\nMOV %s, %s\nPUSH {%s}", register, num.Token.Value, register)
	case parser.UnaryOp:
		unaryOp := leaf.(parser.UnaryOp)
		return compileUnaryOp(unaryOp)
	case parser.BinaryOp:
		binaryOp := leaf.(parser.BinaryOp)
		return compileBinaryOp(binaryOp)
	}
	return ""
}
