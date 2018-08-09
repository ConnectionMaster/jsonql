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
	}

	for i, testCase := range testCases {
		testCaseName := fmt.Sprintf("Test case %d: `%s`", i, testCase.JQL)
		ast, err := Parse(testCase.JQL)
		assert.NoError(t, err, testCaseName+" [parse]")

		val, err := ast.Evaluate(testCase.Data)
		assert.NoError(t, err, testCaseName+fmt.Sprintf(" [evaluate(%v)]", testCase.Data))

		assert.Equal(t, testCase.Expected, val, testCaseName)
		fmt.Printf("pass - %s on %v evaluated to %v\n", testCase.JQL, testCase.Data, val)
	}
}
