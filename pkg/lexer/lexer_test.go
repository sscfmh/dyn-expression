package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	lx := NewLexer("a and b or c == d or e == f or g in [1, 2, 3] or h in ['hello', 'world']")
	tk := lx.Next()
	for tk.ValType != EOF {
		t.Log(tk.ValType, tk.Value)
		tk = lx.Next()
	}
}
