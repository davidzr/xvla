package compiler

import "fmt"

func emitContext(context *Node) {

	child := context.child[0].value
	fmt.Println(child)

}

func Generate(t Node) {
	for _, n := range t.sibling {

		switch n.nodeType {
		case NodeContext:
			emitContext(n)
		}
	}
}
