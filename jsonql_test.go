package jsonql

import (
	"encoding/json"
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
	Data     string
	Expected interface{}
}

func TestParseLiterals(t *testing.T) {
	testCases := []parseTestCase{
		{`null`, "", nil},
		{`true`, "", true},
		{`false`, "", false},
		{`1.25`, "", 1.25},
		{`1.25e2`, "", 125.0},
		{`125e-2`, "", 1.25},
		{`.5`, "", .5},
		{`1`, "", int64(1)},
		{`010`, "", int64(8)},
		{`0xa`, "", int64(10)},
		{`"foo"`, "", "foo"},
		{`"string with \"escape\" characters"`, "", "string with \"escape\" characters"},
		{`'string with \'escape\' characters'`, "", "string with 'escape' characters"},
		{`'peace\u00a0\x26\u00a0war'`, "", "peace & war"},
		{`'\b\f\n\r\t\v'`, "", "\b\f\n\r\t\v"},
		{`"\\"`, "", "\\"},
		{`'\\'`, "", "\\"},
	}

	for i, testCase := range testCases {
		testCaseName := fmt.Sprintf("Test case %d: `%s`", i, testCase.JQL)
		ast, err := Parse(testCase.JQL)
		if !(assert.NoError(t, err, testCaseName+" [parse]") &&
			assert.NotNil(t, ast, testCaseName+" [parse]")) {
			continue
		}

		val, err := ast.Evaluate(nil)
		assert.NoError(t, err, testCaseName+fmt.Sprintf(" [evaluate(nil)]"))

		var passfail = "pass"
		if !assert.Equal(t, testCase.Expected, val, testCaseName) {
			passfail = "fail"
		}
		fmt.Printf("%s - %s evaluated to %v\n", passfail, testCase.JQL, val)
	}
}

func TestParseIdentifiers(t *testing.T) {
	testCases := []parseTestCase{
		{`blah`, `{"blah": "blah"}`, "blah"},
		{`blab`, `{"blah": "blah"}`, nil},
		{`foo.bar`, `{"foo": {"bar": "baz"}}`, "baz"},
		{`foo[1]`, `{"foo": [1, 2, 3]}`, 2.0},
		{`foo[bar]`, `{"foo": ["one", "two", "three"], "bar": 1}`, "two"},
	}

	for i, testCase := range testCases {
		testCaseName := fmt.Sprintf("Test case %d: `%s`", i, testCase.JQL)
		ast, err := Parse(testCase.JQL)
		if !(assert.NoError(t, err, testCaseName+" [parse]") &&
			assert.NotNil(t, ast, testCaseName+" [parse]")) {
			continue
		}

		var data interface{}
		assert.NoError(t, json.Unmarshal([]byte(testCase.Data), &data))
		val, err := ast.Evaluate(data)
		assert.NoError(t, err, testCaseName+fmt.Sprintf(" [evaluate(%v)]", data))

		var passfail = "pass"
		if !assert.Equal(t, testCase.Expected, val, testCaseName) {
			passfail = "fail"
		}
		fmt.Printf("%s - %s on %v evaluated to %v (expected: %v)\n", passfail, testCase.JQL, testCase.Data, val, testCase.Expected)
	}
}
