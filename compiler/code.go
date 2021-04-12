package compiler

import "fmt"

func emitContext(context *Node) {

	child := context.child[0].value
	fmt.Println(child, context.child[0].line)

}

func Generate(t Node) {
	for _, n := range t.Sibling {

		switch n.nodeType {
		case NodeContext:
			emitContext(n)
		}
	}
}
