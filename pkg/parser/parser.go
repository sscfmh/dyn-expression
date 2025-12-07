package parser

import (
	"slices"

	"github.com/sscfmh/dyn-expression/pkg/lexer"
)

type Parser struct {
	Lexer     lexer.Lexer
	CurrToken lexer.Token
}

func NewParser(exp string) *Parser {
	lx := lexer.NewLexer(exp)
	return &Parser{
		Lexer:     *lx,
		CurrToken: *lx.Next(),
	}
}

func (p *Parser) eat(vt lexer.TokenValType) {
	if p.CurrToken.ValType == vt {
		p.CurrToken = *p.Lexer.Next()
	} else {
		panic("Unexpected token")
	}
}

func (p *Parser) factor() Node {
	if p.CurrToken.ValType == lexer.VAR {
		val := p.CurrToken.Value.(string)
		p.eat(lexer.VAR)
		return &Var{VarName: val}
	}
	if p.CurrToken.ValType == lexer.NUM {
		val := p.CurrToken.Value.(int)
		p.eat(lexer.NUM)
		return &Num{Value: val}
	}
	if p.CurrToken.ValType == lexer.STR {
		val := p.CurrToken.Value.(string)
		p.eat(lexer.STR)
		return &Str{Value: val}
	}
	if p.CurrToken.ValType == lexer.LPAREN {
		p.eat(lexer.LPAREN)
		node := p.expr()
		p.eat(lexer.RPAREN)
		return node
	}
	if p.CurrToken.ValType == lexer.LBRACKET {
		p.eat(lexer.LBRACKET)
		var elements []Node
		for {
			elements = append(elements, p.expr())
			if p.CurrToken.ValType == lexer.COMMA {
				p.eat(lexer.COMMA)
			} else {
				break
			}
		}
		p.eat(lexer.RBRACKET)
		return &Arr{Elements: elements}
	}
	return nil
}

var OP_CHARS []lexer.TokenValType = []lexer.TokenValType{lexer.OP_BOOL_EQ, lexer.OP_IN, lexer.OP_AND, lexer.OP_OR}
var FIRST_PIRITY_OPS []lexer.TokenValType = []lexer.TokenValType{lexer.OP_BOOL_EQ, lexer.OP_IN}

func (p *Parser) expr() Node {
	stk := make([]Node, 0)
	stk = append(stk, p.factor())
	for slices.Contains(OP_CHARS, p.CurrToken.ValType) {
		tk := p.CurrToken
		p.eat(tk.ValType)
		var op Node
		switch tk.ValType {
		case lexer.OP_BOOL_EQ:
			op = &BoolEq{}
		case lexer.OP_IN:
			op = &In{}
		case lexer.OP_AND:
			op = &And{}
		case lexer.OP_OR:
			op = &Or{}
		}
		// 在第一优先级的操作
		if slices.Contains(FIRST_PIRITY_OPS, tk.ValType) {
			node := &BinOp{Left: stk[len(stk)-1], Op: op, Right: p.factor()}
			stk = stk[:len(stk)-1]
			stk = append(stk, node)
		} else {
			stk = append(stk, op)
			stk = append(stk, p.factor())
		}
	}
	stk2 := make([]Node, 0)
	opAndFlag := false
	for _, v := range stk {
		if len(stk2) == 0 {
			stk2 = append(stk2, v)
		} else {
			if v.NodeTp() == AND_NODE {
				opAndFlag = true
			} else if v.NodeTp() != OR_NODE {
				if opAndFlag {
					opAndFlag = false
					node := &BinOp{Left: stk2[len(stk2)-1], Op: &And{}, Right: v}
					stk2 = stk2[:len(stk2)-1]
					stk2 = append(stk2, node)
				} else {
					stk2 = append(stk2, v)
				}
			}
		}
	}
	var node Node
	for _, v := range stk2 {
		if node == nil {
			node = v
		} else {
			node = &BinOp{Left: node, Op: &Or{}, Right: v}
		}
	}
	return node
}

func (p *Parser) Parse() Node {
	return p.expr()
}
