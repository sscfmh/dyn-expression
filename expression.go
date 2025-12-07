package expression

import (
	"github.com/sscfmh/dyn-expression/pkg/parser"
)

type NodeVisitor struct {
	p        parser.Parser
	VarTable map[string]interface{}
}

func (nv *NodeVisitor) Visit() bool {
	root := nv.p.Parse()
	return nv.visitNode(root).(bool)
}

func (nv *NodeVisitor) visitNode(node parser.Node) interface{} {
	switch node.NodeTp() {
	case parser.VAR_NODE:
		return nv.VarTable[node.(*parser.Var).VarName]
	case parser.NUM_NODE:
		return node.(*parser.Num).Value
	case parser.STR_NODE:
		return node.(*parser.Str).Value
	case parser.ARR_NODE:
		var elements []interface{}
		for _, v := range node.(*parser.Arr).Elements {
			elements = append(elements, nv.visitNode(v))
		}
		return elements
	case parser.BIN_OP_NODE:
		left := nv.visitNode(node.(*parser.BinOp).Left)
		right := nv.visitNode(node.(*parser.BinOp).Right)
		switch node.(*parser.BinOp).Op.NodeTp() {
		case parser.AND_NODE:
			return left.(bool) && right.(bool)
		case parser.OR_NODE:
			return left.(bool) || right.(bool)
		case parser.BOOL_EQ_NODE:
			return left == right
		case parser.IN_NODE:
			for _, v := range right.([]interface{}) {
				if left == v {
					return true
				}
			}
			return false
		}
	}
	return nil
}

func Execute(exp string, varTable map[string]interface{}) bool {
	p := parser.NewParser(exp)
	nv := NodeVisitor{p: *p, VarTable: varTable}
	return nv.Visit()
}
