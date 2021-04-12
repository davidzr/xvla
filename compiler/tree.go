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
	child    []*Node
	Sibling  []*Node
	line     int
}

func NewNode(t NodeType, line int) *Node {
	node := Node{
		nodeType: t,
		child:    []*Node{},
		value:    "",
		Sibling:  []*Node{},
		line:     line,
	}
	return &node
}
