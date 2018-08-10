package ast

import (
	"fmt"

	"math"
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
		indexFloat, ok := indexVal.(float64)
		if !ok {
			return nil, fmt.Errorf("Index expression yielded %v, not an integer", indexVal)
		}
		indexInt = int64(indexFloat)
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

type NegateNode struct {
	Expr Expr
}

func Negative(val interface{}) (Expr, error) {
	return NegateNode{val.(Expr)}, nil
}

func (nn NegateNode) Evaluate(val interface{}) (interface{}, error) {
	val, err := nn.Expr.Evaluate(val)
	if err != nil {
		return nil, err
	}
	switch v := val.(type) {
	case nil:
		return nil, nil
	case int64:
		return -v, nil
	case int:
		return -v, nil
	case float64:
		return -v, nil
	default:
		// type error... nil for now
		return nil, nil
	}
}

type NotNode struct {
	Expr Expr
}

func Not(val interface{}) (Expr, error) {
	return NotNode{val.(Expr)}, nil
}

func (nn NotNode) Evaluate(val interface{}) (interface{}, error) {
	val, err := nn.Expr.Evaluate(val)
	if err != nil {
		return nil, err
	}
	switch v := val.(type) {
	case nil:
		return nil, nil
	case bool:
		return !v, nil
	case int64:
		return (v == 0), nil
	case int:
		return (v == 0), nil
	case float64:
		// equal to zero is strong for a float but hey
		return (v == 0), nil
	default:
		// type error... nil for now
		return nil, nil
	}
}

type ExpNode struct {
	Base Expr
	Pow  Expr
}

func Exp(base, pow interface{}) (Expr, error) {
	return ExpNode{base.(Expr), pow.(Expr)}, nil
}

func (en ExpNode) Evaluate(val interface{}) (interface{}, error) {
	base, err := en.Base.Evaluate(val)
	if err != nil {
		return nil, err
	}
	pow, err := en.Pow.Evaluate(val)
	if err != nil {
		return nil, err
	}
	switch p := pow.(type) {
	case nil:
		return nil, nil
	case int64:
		switch b := base.(type) {
		case int64:
			pow := math.Pow(float64(b), float64(p))
			if float64(int64(pow)) != pow {
				return pow, nil
			}
			return int64(pow), nil
		case float64:
			return math.Pow(b, float64(p)), nil
		default:
			return nil, nil
		}
	case float64:
		switch b := base.(type) {
		case int64:
			return math.Pow(float64(b), p), nil
		case float64:
			return math.Pow(b, p), nil
		default:
			return nil, nil
		}
	default:
		// type error... nil for now
		return nil, nil
	}
}
