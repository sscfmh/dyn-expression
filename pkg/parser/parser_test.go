package parser

import (
	"encoding/json"
	"testing"

	"github.com/sscfmh/dyn-expression/pkg/lexer"
)

func TestParser(t *testing.T) {
	lx := lexer.NewLexer("a and b or c == d or e == f or g in [1, 2, 3] or h in ['hello', 'world']")
	p := Parser{
		Lexer:     *lx,
		CurrToken: *lx.Next(),
	}
	node := p.Parse()
	dump, _ := json.MarshalIndent(node, "--", "    ")
	t.Log(string(dump))
}
