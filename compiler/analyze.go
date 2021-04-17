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
		case nodeVariable:
			_, ok := symtab[n.name]
			value := n.child[0].value
			if !ok {
				symtab[n.name] = Resource{
					typeo: "variable",
					value: value,
				}
			} else {
				typeError(n, "already declared.")
			}

		case nodeRule:
			_, ok := symtab[n.name]

			if !ok {
				symtab[n.name] = Resource{
					typeo: "rule",
					value: "value",
				}
			} else {
				typeError(n, "already declared.")
			}
		case nodeNamespace:
			_, ok := symtab[n.name]
			if !ok {
				symtab[n.name] = Resource{
					typeo: "namespace",
				}
			} else {
				typeError(n, "already declared.")
			}
		case nodeContextBody, nodeRuleBody, nodeContext:
			analyze(n.child)
		}

	}
}

func typeCheck(nodes []*Node) {
	for _, n := range nodes {
		switch n.nodeType {
		case nodeContext:
			if n.child[0].nodeType == nodeReference {
				name := n.child[0].name[1:]
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
		case nodeApply:
			name := n.child[0].name[1:]
			tree, ok := symtab[name]
			if ok {
				if tree.typeo != "rule" {
					typeError(n, "reference must be a rule")
				}
			} else {
				typeError(n, "rule not found.")
			}
		case nodeRule, nodeContextBody, nodeRuleBody:
			typeCheck(n.child)
		}
	}
}
