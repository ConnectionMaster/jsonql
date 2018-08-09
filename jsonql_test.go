package jsonql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {

	jsonString := `
[
  {
    "name": "elgs",
    "gender": "m",
    "skills": [
      "Golang",
      "Java",
      "C"
    ]
  },
  {
    "name": "enny",
    "gender": "f",
    "skills": [
      "IC",
      "Electric design",
      "Verification"
    ]
  },
  {
    "name": "sam",
    "gender": "m",
    "skills": [
      "Eating",
      "Sleeping",
      "Crawling"
    ]
  }
]
`
	parser, err := NewStringQuery(jsonString)
	if err != nil {
		t.Error(err)
	}

	var pass = []struct {
		in string
		ex int
	}{
		{"name='elgs'", 1},
		{"gender='f'", 1},
		{"skills.[1]='Sleeping'", 1},
	}
	var fail = []struct {
		in string
		ex interface{}
	}{}
	for _, v := range pass {
		result, err := parser.Query(v.in)
		if err != nil {
			t.Error(v.in, err)
		}
		fmt.Println(v.in, result)
		//		if v.ex != result {
		//			t.Error("Expected:", v.ex, "actual:", result)
		//		}
	}
	for range fail {

	}
}

type parseTestCase struct {
	JQL      string
	Data     interface{}
	Expected interface{}
}

func TestParseLiterals(t *testing.T) {
	testCases := []parseTestCase{
		{`null`, nil, nil},
		{`true`, nil, true},
		{`false`, nil, false},
		{`1.25`, nil, 1.25},
		{`1.25e2`, nil, 125.0},
		{`125e-2`, nil, 1.25},
		{`.5`, nil, .5},
		{`1`, nil, int64(1)},
		{`010`, nil, int64(8)},
		{`0xa`, nil, int64(10)},
		{`"foo"`, nil, "foo"},
		{`"string with \"escape\" characters"`, nil, "string with \"escape\" characters"},
		{`'string with \'escape\' characters'`, nil, "string with 'escape' characters"},
		{`'peace\u00a0\x26\u00a0war'`, nil, "peace & war"},
		{`'\b\f\n\r\t\v'`, nil, "\b\f\n\r\t\v"},
		{`"\\"`, nil, "\\"},
		{`'\\'`, nil, "\\"},
	}

	for i, testCase := range testCases {
		testCaseName := fmt.Sprintf("Test case %d: `%s`", i, testCase.JQL)
		ast, err := Parse(testCase.JQL)
		if !(assert.NoError(t, err, testCaseName+" [parse]") &&
			assert.NotNil(t, ast, testCaseName+" [parse]")) {
			continue
		}

		val, err := ast.Evaluate(testCase.Data)
		assert.NoError(t, err, testCaseName+fmt.Sprintf(" [evaluate(%v)]", testCase.Data))

		var passfail = "pass"
		if !assert.Equal(t, testCase.Expected, val, testCaseName) {
			passfail = "fail"
		}
		fmt.Printf("%s - %s on %v evaluated to %v (expected: %v)\n", passfail, testCase.JQL, testCase.Data, val, testCase.Expected)
	}
}
