package ast

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Expr interface {
	Evaluate(data interface{}) (interface{}, error)
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

type BinaryOpNode struct {
	Op       OpType
	Operands [2]Expr
}

func (bon BinaryOpNode) Evaluate(val interface{}) (interface{}, error) {
	left, err := bon.Operands[0].Evaluate(val)
	if err != nil {
		return nil, err
	}
	right, err := bon.Operands[1].Evaluate(val)
	if err != nil {
		return nil, err
	}
	switch left := left.(type) {
	case nil:
		return nil, nil
	case int64:
		switch right := right.(type) {
		case int64:
			return bon.Op.Int(left, right)
		case float64:
			return bon.Op.Float(float64(left), right)
		default:
			return nil, nil
		}
	case float64:
		switch right := right.(type) {
		case int64:
			return bon.Op.Float(left, float64(right))
		case float64:
			return bon.Op.Float(left, right)
		default:
			return nil, nil
		}
	default:
		// type error... nil for now
		return nil, nil
	}
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

func (op OpType) Int(left, right int64) (interface{}, error) {
	switch op {
	case OpAdd:
		return left + right, nil
	case OpSub:
		return left - right, nil
	case OpMul:
		return left * right, nil
	case OpDiv:
		if right == 0 {
			return nil, nil
		}
		quot := left / right
		if quot*right == left {
			return quot, nil
		}
		return float64(left) / float64(right), nil
	case OpMod:
		if right == 0 {
			return nil, nil
		}
		return left % right, nil
	case OpExp:
		// cheat :)
		pow := math.Pow(float64(left), float64(right))
		if float64(int64(pow)) != pow {
			return pow, nil
		}
		return int64(pow), nil
	default:
		return nil, nil
	}
}

func (op OpType) Float(left, right float64) (interface{}, error) {
	switch op {
	case OpAdd:
		return left + right, nil
	case OpSub:
		return left - right, nil
	case OpMul:
		return left * right, nil
	case OpDiv:
		if right == 0 {
			return nil, nil
		}
		return left / right, nil
	case OpMod:
		return math.Mod(left, right), nil
	case OpExp:
		return math.Pow(left, right), nil
	default:
		return nil, nil
	}
}

func Mul(multiplicand, multiplier interface{}) (Expr, error) {
	return BinaryOpNode{OpMul, [2]Expr{multiplicand.(Expr), multiplier.(Expr)}}, nil
}

func Div(dividend, divisor interface{}) (Expr, error) {
	return BinaryOpNode{OpDiv, [2]Expr{dividend.(Expr), divisor.(Expr)}}, nil
}

func Mod(dividend, modulus interface{}) (Expr, error) {
	return BinaryOpNode{OpMod, [2]Expr{dividend.(Expr), modulus.(Expr)}}, nil
}

func Add(augend, addend interface{}) (Expr, error) {
	return BinaryOpNode{OpAdd, [2]Expr{augend.(Expr), addend.(Expr)}}, nil
}

func Sub(minuend, subtrahend interface{}) (Expr, error) {
	return BinaryOpNode{OpSub, [2]Expr{minuend.(Expr), subtrahend.(Expr)}}, nil
}

func Exp(base, exponent interface{}) (Expr, error) {
	return BinaryOpNode{OpExp, [2]Expr{base.(Expr), exponent.(Expr)}}, nil
}

type RegexpOp struct {
	Match   Expr
	Pattern regexp.Regexp
	Inverse bool
}

func RegexpMatch(match, pattern interface{}) (Expr, error) {
	return NewRegexpOp(match, pattern.(string), false)
}

func RegexpNegMatch(match, pattern interface{}) (Expr, error) {
	return NewRegexpOp(match, pattern.(string), true)
}

func NewRegexpOp(match interface{}, pattern string, inverse bool) (RegexpOp, error) {
	var regexpOp = RegexpOp{
		Inverse: inverse,
	}
	if matchStr, ok := match.(string); ok {
		regexpOp.Match = LiteralNode{matchStr}
	} else {
		regexpOp.Match = match.(Expr)
	}
	patternRegexp, err := regexp.Compile(pattern)
	if err != nil {
		return regexpOp, err
	}
	regexpOp.Pattern = *patternRegexp
	return regexpOp, nil
}

func (ro RegexpOp) Evaluate(val interface{}) (interface{}, error) {
	value, err := ro.Match.Evaluate(val)
	if err != nil {
		return nil, err
	}
	switch value := value.(type) {
	case []byte:
		return xor(ro.Pattern.Match(value), ro.Inverse), nil
	case string:
		return xor(ro.Pattern.MatchString(value), ro.Inverse), nil
	case int64:
		return xor(ro.Pattern.MatchString(strconv.FormatInt(value, 10)), ro.Inverse), nil
	case float64:
		return xor(ro.Pattern.MatchString(strconv.FormatFloat(value, 'f', -1, 64)), ro.Inverse), nil
	default:
		return nil, nil
	}
}

func xor(a, b bool) bool {
	return ((a && !b) || (!a && b))
}
