package compiler

func Run(source string) {
	tree := parse(source)
	analyze(tree.sibling)
	typeCheck(tree.sibling)

}
