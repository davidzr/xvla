package compiler

type Resource struct {
	value string
	line  int
	typeo string
}

var symtab = make(map[string]Resource)

func Analyze(nodes []*Node) {
	for _, n := range nodes {

		switch n.nodeType {
		case NodeVariable:
			_, ok := symtab[n.name]
			if !ok {
				symtab[n.name] = Resource{
					typeo: "variable",
				}
			}
		case NodeRule:
			_, ok := symtab[n.name]
			if !ok {
				symtab[n.name] = Resource{
					typeo: "rule",
				}
			}
		case NodeNamespace:
			_, ok := symtab[n.name]
			if !ok {
				symtab[n.name] = Resource{
					typeo: "namespace",
				}
			}
		case NodeContextBody:
			Analyze(n.child)
		case NodeRuleBody:
			Analyze(n.child)
		}

	}
}
