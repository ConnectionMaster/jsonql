package ast

type Expr interface {
	Evaluate(data interface{}) (interface{}, error)
}

type OpType int

const (
	OpAdd OpType = iota
	OpSub
	OpMul
	OpDiv
	OpMod
	OpExp
	OpAnd
	OpOr
)

type ExprOp struct {
	Op       OpType
	Operands []Expr
}

// LiteralNode is an AST node for an immediate value
type LiteralNode struct {
	val interface{}
}

func (LN LiteralNode) Evaluate(interface{}) (interface{}, error) {
	return LN.val, nil
}

func Literal(val interface{}) (interface{}, error) {
	return LiteralNode{val}, nil
}
