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

			name := n.child[0].value
			_, ok := symtab[name]

			if !ok {
				symtab[name] = Resource{
					value: n.child[1].value,
					typeo: "string",
				}
			} else {
				panic("Already declared variable")
			}
		case NodeIdentifier:
			name := n.value[1:]

			_, ok := symtab[name]

			if !ok {
				panic("Not declared variable")
			}

		case NodeRule:
			Analyze(n.child)
		}
	}
}
