package ast

import (
	"fmt"
)

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

func Literal(val interface{}) (Expr, error) {
	return LiteralNode{val}, nil
}

type ObjectKeyNode struct {
	Key    string
	Parent Expr
}

func ObjectKey(val interface{}) (Expr, error) {
	return ObjectKeyNode{val.(string), nil}, nil
}

func SelectKey(base Expr, val interface{}) (Expr, error) {
	return ObjectKeyNode{val.(string), base}, nil
}

func (okn ObjectKeyNode) Evaluate(val interface{}) (interface{}, error) {
	if okn.Parent != nil {
		var err error
		val, err = okn.Parent.Evaluate(val)
		if err != nil {
			// TODO: better wrapping of contextual errors
			return nil, err
		}
	}
	if mapVal, ok := val.(map[string]interface{}); ok {
		return mapVal[okn.Key], nil
	}
	return nil, nil
}

type IndexNode struct {
	Index  Expr
	Parent Expr
}

func Index(base Expr, index Expr) (Expr, error) {
	return IndexNode{index, base}, nil
}

func (in IndexNode) Evaluate(val interface{}) (interface{}, error) {
	indexVal, err := in.Index.Evaluate(val)
	if err != nil {
		return nil, fmt.Errorf("Error evaluating index expression: %s", err.Error())
	}
	if indexVal == nil {
		return nil, nil
	}
	indexInt, ok := indexVal.(int64)
	if !ok {
		return nil, fmt.Errorf("Index expression yielded %v, not an integer", indexVal)
	}
	val, err = in.Parent.Evaluate(val)
	if arrVal, ok := val.([]interface{}); ok {
		if abs(indexInt) >= int64(len(arrVal)) {
			return nil, nil
		}
		if indexInt < 0 {
			return arrVal[len(arrVal)-int(indexInt)], nil
		}
		return arrVal[int(indexInt)], nil
	}
	return nil, nil
}

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
