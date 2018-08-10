package jsonql_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/teslamotors/jsonql/lexer"
	"github.com/teslamotors/jsonql/token"
)

func TestTokenize(t *testing.T) {
	var pass = []struct {
		in string
		ex []string
	}{
		{"-1 + 2", []string{"-", "1", "+", "2"}},
		{"-(1+2)", []string{"-", "(", "1", "+", "2", ")"}},
		// {"+1+2", []string{"+1", "+", "2"}},
		{"1+2+(3*4)", []string{"1", "+", "2", "+", "(", "3", "*", "4", ")"}},
		{"1+2+(3*4)^34", []string{"1", "+", "2", "+", "(", "3", "*", "4", ")", "^", "34"}},
		{"'123  456' 789", []string{"'123  456'", "789"}},
		{`123 "456  789"`, []string{"123", "\"456  789\""}},
		{`123 "456  '''789"`, []string{"123", "\"456  '''789\""}},
	}
	for _, v := range pass {
		fmt.Printf("Parsing: %#v\n", v.in)
		scanner := lexer.NewLexer([]byte(v.in))
		var got []string
		for tok := scanner.Scan(); tok.Type != token.EOF; tok = scanner.Scan() {
			got = append(got, string(tok.Lit))
		}
		assert.Equal(t, v.ex, got, fmt.Sprintf("Tokenized: %#v", v.in))
	}
}
