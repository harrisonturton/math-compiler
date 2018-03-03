package parser

import (
	"fmt"
	"../token"
)

type Expr interface {
}

type Number struct {
	Token token.Token
}

type UnaryOp struct {
	Op    token.Token
	Right Expr
}

type BinaryOp struct {
	Left  Expr
	Op    token.Token
	Right Expr
}

func (n Number) String() string {
	return fmt.Sprintf("%s", n.Token.Value)
}

func (b BinaryOp) String() string {
	//return fmt.Sprintf("(%s %s %s)", b.Left, b.Op, b.Right)
	return fmt.Sprintf("(%s %s %s)", b.Op, b.Left, b.Right)
}

func (u UnaryOp) String() string {
	return fmt.Sprintf("(%s %s)", u.Op, u.Right)
}
