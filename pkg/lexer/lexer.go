package lexer

import (
	"strconv"
	"unicode"
)

type Lexer struct {
	Text     []rune
	Pos      int
	CurrChar *rune
}

func NewLexer(content string) *Lexer {
	text := make([]rune, 0)
	for _, c := range content {
		text = append(text, c)
	}
	return &Lexer{
		Text:     text,
		Pos:      0,
		CurrChar: &text[0],
	}
}

func (lx *Lexer) forward() {
	lx.Pos += 1
	if lx.Pos < len(lx.Text) {
		lx.CurrChar = &lx.Text[lx.Pos]
	} else {
		lx.CurrChar = nil
	}
}

func (lx *Lexer) skipSpace() {
	for lx.CurrChar != nil && unicode.IsSpace(*lx.CurrChar) {
		lx.forward()
	}
}

func (lx *Lexer) id() *Token {
	s := make([]rune, 0)
	for lx.CurrChar != nil && (unicode.IsLetter(*lx.CurrChar) || unicode.IsNumber(*lx.CurrChar)) {
		s = append(s, *lx.CurrChar)
		lx.forward()
	}
	val := string(s)
	if RESERVED_WORDS[val] != "" {
		return &Token{
			ValType: RESERVED_WORDS[val],
			Value:   nil,
		}
	}
	return &Token{
		ValType: VAR,
		Value:   val,
	}
}

func (lx *Lexer) str() *Token {
	lx.forward()
	s := make([]rune, 0)
	for *lx.CurrChar != '\'' {
		s = append(s, *lx.CurrChar)
		lx.forward()
	}
	lx.forward()
	return &Token{
		ValType: STR,
		Value:   string(s),
	}
}

func (lx *Lexer) num() *Token {
	s := make([]rune, 0)
	for lx.CurrChar != nil && unicode.IsNumber(*lx.CurrChar) {
		s = append(s, *lx.CurrChar)
		lx.forward()
	}
	num, _ := strconv.Atoi(string(s))
	return &Token{
		ValType: NUM,
		Value:   num,
	}
}

func (lx *Lexer) Next() *Token {
	for lx.CurrChar != nil {
		if unicode.IsLetter(*lx.CurrChar) {
			return lx.id()
		}
		if unicode.IsSpace(*lx.CurrChar) {
			lx.skipSpace()
		}
		if *lx.CurrChar == '\'' {
			return lx.str()
		}
		if *lx.CurrChar == '(' {
			lx.forward()
			return &Token{
				ValType: LPAREN,
				Value:   nil,
			}
		}
		if *lx.CurrChar == ')' {
			lx.forward()
			return &Token{
				ValType: RPAREN,
				Value:   nil,
			}
		}
		if *lx.CurrChar == '=' {
			lx.forward()
			lx.forward()
			return &Token{
				ValType: OP_BOOL_EQ,
				Value:   nil,
			}
		}
		if unicode.IsNumber(*lx.CurrChar) {
			return lx.num()
		}
		if *lx.CurrChar == '[' {
			lx.forward()
			return &Token{
				ValType: LBRACKET,
				Value:   nil,
			}
		}
		if *lx.CurrChar == ']' {
			lx.forward()
			return &Token{
				ValType: RBRACKET,
				Value:   nil,
			}
		}
		if *lx.CurrChar == ',' {
			lx.forward()
			return &Token{
				ValType: COMMA,
				Value:   nil,
			}
		}
	}
	return &Token{
		ValType: EOF,
		Value:   nil,
	}
}
