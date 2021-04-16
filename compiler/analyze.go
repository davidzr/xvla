package compiler

import "fmt"

type Resource struct {
	value string
	line  int
	typeo string
}

var symtab = make(map[string]Resource)

func typeError(t *Node, message string) {
	errorStr := fmt.Sprintf("Type error at: %d, %s", t.line, message)
	panic(errorStr)
}

func analyze(nodes []*Node) {
	for _, n := range nodes {

		switch n.nodeType {
		case NodeVariable:
			_, ok := symtab[n.name]
			if !ok {
				symtab[n.name] = Resource{
					typeo: "variable",
				}
			} else {
				typeError(n, "already declared.")
			}

		case NodeRule:
			_, ok := symtab[n.name]
			if !ok {
				symtab[n.name] = Resource{
					typeo: "rule",
				}
			} else {
				typeError(n, "already declared.")
			}
		case NodeNamespace:
			_, ok := symtab[n.name]
			if !ok {
				symtab[n.name] = Resource{
					typeo: "namespace",
				}
			} else {
				typeError(n, "already declared.")
			}
		case NodeContextBody:
			analyze(n.child)
		case NodeRuleBody:
			analyze(n.child)
		}

	}
}

func typeCheck(nodes []*Node) {
	for _, n := range nodes {
		switch n.nodeType {
		case NodeContext:
			if n.child[0].nodeType == NodeReference {
				name := n.child[0].name[1:]
				fmt.Println(name)
				tree, ok := symtab[name]
				if ok {
					if tree.typeo != "variable" {
						typeError(n, "path is not a variable or string.")
					}
				} else {
					typeError(n, "variable not found")
				}
			}
			typeCheck(n.child)
		case NodeApply:
			name := n.child[0].name[1:]
			tree, ok := symtab[name]
			if ok {
				if tree.typeo != "rule" {
					typeError(n, "identifier mus be a rule")
				}
			} else {
				typeError(n, "rule not found.")
			}
		case NodeRule, NodeContextBody, NodeRuleBody:
			typeCheck(n.child)
		}
	}
}
