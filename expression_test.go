package expression

import (
	"testing"
)

func TestExecute(t *testing.T) {
	varTable := map[string]interface{}{
		"a": true,
		"b": false,
		"c": 1,
		"d": 2,
		"e": "hello",
		"f": "world",
		"g": 1,
		"h": []interface{}{"hello", "world"},
	}
	ret := Execute("a and b or c == d or e == f or g in [1, 2, 3] or h in ['hello', 'world']", varTable)
	t.Log("result: ", ret)
}
