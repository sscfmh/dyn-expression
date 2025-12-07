package lexer

type TokenValType string

const (
	VAR      TokenValType = "VAR"
	LPAREN   TokenValType = "LPAREN"
	RPAREN   TokenValType = "RPAREN"
	LBRACKET TokenValType = "LBRACKET"
	RBRACKET TokenValType = "RBRACKET"
	COMMA    TokenValType = "COMMA"
	EOF      TokenValType = "EOF"

	OP_IN      TokenValType = "OP_IN"
	OP_AND     TokenValType = "OP_AND"
	OP_OR      TokenValType = "OP_OR"
	OP_BOOL_EQ TokenValType = "OP_BOOL_EQ"

	ARRAY TokenValType = "ARRAY"
	STR   TokenValType = "STR"
	NUM   TokenValType = "NUM"
)

var RESERVED_WORDS = map[string]TokenValType{
	"in":  OP_IN,
	"and": OP_AND,
	"or":  OP_OR,
}

type Token struct {
	ValType TokenValType
	Value   interface{}
}
