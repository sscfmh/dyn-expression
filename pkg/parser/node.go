package parser

type NodeType string

const (
	VAR_NODE     NodeType = "VAR_NODE"
	STR_NODE     NodeType = "STR_NODE"
	NUM_NODE     NodeType = "NUM_NODE"
	ARR_NODE     NodeType = "ARR_NODE"
	BIN_OP_NODE  NodeType = "BIN_OP_NODE"
	OR_NODE      NodeType = "OR_NODE"
	AND_NODE     NodeType = "AND_NODE"
	BOOL_EQ_NODE NodeType = "BOOL_EQ_NODE"
	IN_NODE      NodeType = "IN_NODE"
)

type Node interface {
	NodeTp() NodeType
}

type Var struct {
	VarName string
}

func (v *Var) NodeTp() NodeType {
	return VAR_NODE
}

type Str struct {
	Value string
}

func (s *Str) NodeTp() NodeType {
	return STR_NODE
}

type Num struct {
	Value int
}

func (n *Num) NodeTp() NodeType {
	return NUM_NODE
}

type Arr struct {
	Elements []Node
}

func (a *Arr) NodeTp() NodeType {
	return ARR_NODE
}

type BinOp struct {
	Left  Node
	Op    Node
	Right Node
}

func (b *BinOp) NodeTp() NodeType {
	return BIN_OP_NODE
}

type Or struct{}

func (o *Or) NodeTp() NodeType {
	return OR_NODE
}

type And struct{}

func (a *And) NodeTp() NodeType {
	return AND_NODE
}

type BoolEq struct{}

func (b *BoolEq) NodeTp() NodeType {
	return BOOL_EQ_NODE
}

type In struct{}

func (i *In) NodeTp() NodeType {
	return IN_NODE
}
