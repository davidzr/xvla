package compiler

type NodeType int

const (
	NodeContext = iota
	NodeRule
	NodeAssert
	NodeVariable
	NodeNamespace
	NodeReturn
	NodeString
	NodeReference
	NodeIdentifier
	NodeContextBody
	NodeRuleBody
	NodeApply
)

type Node struct {
	nodeType NodeType
	value    string
	name     string
	child    []*Node
	sibling  []*Node
	line     int
}

func NewNode(t NodeType, line int) *Node {
	node := Node{
		nodeType: t,
		child:    []*Node{},
		value:    "",
		sibling:  []*Node{},
		line:     line,
	}
	return &node
}
