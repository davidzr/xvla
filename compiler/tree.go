package compiler

type NodeType int

const (
	nodeContext = iota
	nodeRule
	nodeAssert
	nodeVariable
	nodeNamespace
	nodeReturn
	nodeLiteral
	nodeReference
	nodeIdentifier
	nodeContextBody
	nodeRuleBody
	nodeApply
	nodeProgram
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
