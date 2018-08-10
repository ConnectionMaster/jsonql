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
		testCaseName := fmt.Sprintf("Parse Literal case %d: `%s`", i, testCase.JQL)
		assertTestCase(t, testCase, testCaseName)
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
		testCaseName := fmt.Sprintf("Parse Identifier case %d: `%s`", i, testCase.JQL)
		assertTestCase(t, testCase, testCaseName)
	}
}

func TestUnaryExpressions(t *testing.T) {
	testCases := []parseTestCase{
		{`!null`, ``, nil},
		{`!true`, ``, false},
		{`!false`, ``, true},
		{`!0`, ``, true},
		{`!1`, ``, false},
		{`!0.0`, ``, true},
		{`!1.0`, ``, false},
		{`!blah`, `{"blah": true}`, false},
		{`!blah`, `{"blah": false}`, true},
		{`!!blah`, `{"blah": 0.0}`, false},
		{`!blah.blah`, `{"blah": {"blah": false}}`, true},
		{`!blah.baa`, `{"blah": {"blah": false}}`, nil},
		{`!!blah.baa`, `{"blah": {"blah": false}}`, nil},
		{`-null`, ``, nil},
		{`-0`, ``, int64(0)},
		{`-1`, ``, int64(-1)},
		{`-0.0`, ``, 0.0},
		{`-1.0`, ``, -1.0},
		{`-blah`, `{"blah": 42.0}`, -42.0},
	}

	for i, testCase := range testCases {
		testCaseName := fmt.Sprintf("Parse Unary Expression case %d: `%s`", i, testCase.JQL)
		assertTestCase(t, testCase, testCaseName)
	}
}

func TestBEDMASExpr(t *testing.T) {
	testCases := []parseTestCase{
		{`2^0`, ``, int64(1)},
		{`2^10`, ``, int64(1024)},
		{`2.0^0`, ``, 1.0},
		{`4^0.5`, ``, 2.0},
		{`256.0^0.25`, ``, 4.0},
		{`2^-2`, ``, 0.25},
		{`2*2`, ``, int64(4)},
		{`2^2*4`, ``, int64(16)},
		{`4*2^-2`, ``, 1.0},
		{`1/8`, ``, 0.125},
		{`1/2^-1`, ``, 2.0},
		{`8%5`, ``, int64(3)},
		{`8%5.0`, ``, 3.0},
		{`10-2*4`, ``, int64(2)},
		{`10-2-4`, ``, int64(4)},
		{`10*2+1/4`, ``, 20.25},
	}

	for i, testCase := range testCases {
		testCaseName := fmt.Sprintf("Parse BEDMAS case %d: `%s`", i, testCase.JQL)
		assertTestCase(t, testCase, testCaseName)
	}
}

func TestRegexpExpr(t *testing.T) {
	testCases := []parseTestCase{
		{`"Hello" ~= "He"`, ``, true},
		{`"Hello" ~= "^el"`, ``, false},
		{`"Hello" ~= ".*el"`, ``, true},
		{`"Hello" !~= "He"`, ``, false},
		{`"Hello" !~= "^el"`, ``, true},
		{`"Hello" !~= ".*el"`, ``, false},
		{`foo.bar ~= "\\d+(\\.\\d*)?$"`, `{"foo": {"bar": "1.23"}}`, true},
		{`foo.bar ~= "\\d+(\\.\\d*)?$"`, `{"foo": {"bar": 22}}`, true},
		{`foo.bar ~= "\\d+(\\.\\d*)?$"`, `{"foo": {"bar": 22.1}}`, true},
		{`foo.bar ~= "\\d+(\\.\\d*)?$"`, `{"foo": {"bar": "blah"}}`, false},
	}

	for i, testCase := range testCases {
		testCaseName := fmt.Sprintf("Parse Regexp case %d: `%s`", i, testCase.JQL)
		assertTestCase(t, testCase, testCaseName)
	}
}

func TestCompareExpr(t *testing.T) {
	testCases := []parseTestCase{
		{`"Hello" = "Hello"`, ``, true},
		{`"Hello" != "Hello"`, ``, false},
		{`message.body = "Hello"`, `{"message":{"body":"Blah"}}`, false},
		{`message.body != "Blah"`, `{"message":{"body":"Hello"}}`, true},
		{`message.body = "Hello"`, `{"message":{"head":"Blah"}}`, nil},
		{`message.body != "Blah"`, `{"message":{"head":"Hello"}}`, nil},
		{`1 < 2`, ``, true},
		{`1 > 2`, ``, false},
		{`3 <= 2`, ``, false},
		{`3 >= 2`, ``, true},
		{`2.0 >= 2`, ``, true},
		{`2 = 2`, ``, true},
		{`null is not defined`, ``, true},
		{`null isnot defined`, ``, true},
		{`null is defined`, ``, false},
		{`null is not null`, ``, false},
		{`null isnot null`, ``, false},
		{`null is null`, ``, true},
		{`message.body is not defined`, `{"message":{"body":"Blah"}}`, false},
		{`message.body is defined`, `{"message":{"body":"Hello"}}`, true},
		{`message.body is not null`, `{"message":{"head":"Blah"}}`, false},
		{`message.body is null`, `{"message":{"head":"Hello"}}`, true},
	}

	for i, testCase := range testCases {
		testCaseName := fmt.Sprintf("Comparison case %d: `%s`", i, testCase.JQL)
		assertTestCase(t, testCase, testCaseName)
	}
}

func assertTestCase(t *testing.T, testCase parseTestCase, testCaseName string) {
	ast, err := Parse(testCase.JQL)
	if !(assert.NoError(t, err, testCaseName+" [parse]") &&
		assert.NotNil(t, ast, testCaseName+" [parse]")) {
		return
	}

	var data interface{}
	if len(testCase.Data) > 0 {
		if !assert.NoError(t, json.Unmarshal([]byte(testCase.Data), &data)) {
			return
		}
	}
	val, err := ast.Evaluate(data)
	if !assert.NoError(t, err, testCaseName+fmt.Sprintf(" [evaluate(%v)]", data)) {
		return
	}
	var passfail = "pass"
	if !assert.Equal(t, testCase.Expected, val, testCaseName) {
		passfail = "fail"
	}
	fmt.Printf("%s - %s on %v evaluated to %v (expected: %v)\n",
		passfail, testCase.JQL, testCase.Data, val, testCase.Expected)
}
