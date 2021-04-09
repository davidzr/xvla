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
	sibling  []*Node
}

func NewNode(t NodeType) *Node {
	node := Node{
		nodeType: t,
		child:    []*Node{},
		value:    "",
		sibling:  []*Node{},
	}
	return &node
}
