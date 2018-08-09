package jsonql

import (
	"encoding/json"
	"fmt"

	"github.com/teslamotors/jsonql/ast"
	"github.com/teslamotors/jsonql/lexer"
	"github.com/teslamotors/jsonql/parser"
)

// JSONQL - JSON Query Lang struct encapsulating the JSON data.
type JSONQL struct {
	Data interface{}
}

// NewStringQuery - creates a new &JSONQL from raw JSON string
func NewStringQuery(jsonString string) (*JSONQL, error) {
	var data = new(interface{})
	err := json.Unmarshal([]byte(jsonString), data)
	if err != nil {
		return nil, err
	}
	return &JSONQL{*data}, nil
}

// NewQuery - creates a new &JSONQL from an array of interface{} or a map of [string]interface{}
func NewQuery(jsonObject interface{}) *JSONQL {
	return &JSONQL{jsonObject}
}

// Query - queries against the JSON using the conditions specified in the where stirng.
func Parse(expr string) (ast.Expr, error) {
	lex := lexer.NewLexer([]byte(expr))
	p := parser.NewParser()
	st, err := p.Parse(lex)
	if err != nil {
		return nil, err
	}
	astExpr, ok := st.(ast.Expr)
	if !ok {
		return nil, fmt.Errorf("parsing JSONQL returned a %T", st)
	}
	return astExpr, nil
}

// Query is a JSON QL query
func (thisJSONQL *JSONQL) Query(where string) (interface{}, error) {
	ast, err := Parse(where)
	if err != nil {
		return nil, err
	}
	return thisJSONQL.QueryExpr(ast)
}

// Query - queries against the JSON using the conditions specified in the where stirng.
func (thisJSONQL *JSONQL) QueryExpr(expr ast.Expr) (interface{}, error) {
	switch v := thisJSONQL.Data.(type) {
	case []interface{}:
		ret := []interface{}{}
		for _, obj := range v {
			val, err := expr.Evaluate(obj)
			if err != nil {
				return nil, err
			}
			switch val {
			case nil, false:
			default:
				ret = append(ret, obj)
			}
		}
		return ret, nil
	case map[string]interface{}:
		val, err := expr.Evaluate(v)
		if err != nil {
			return nil, err
		}
		switch val {
		case nil, false, int(0), "":
			return nil, nil
		default:
			return v, nil
		}
	default:
		return nil, fmt.Errorf("Failed to parse input data.")
	}
}
