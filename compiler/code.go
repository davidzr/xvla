package compiler

import (
	"encoding/xml"
	"fmt"
	"html"
)

type schema struct {
	Ns      []ns      `xml:"ns"`
	Let     []let     `xml:"let"`
	Pattern []pattern `xml:"pattern"`
}

type pattern struct {
	Rule []rule `xml:"rule"`
	Name string `xml:"name,attr"`
}
type rule struct {
	Context string   `xml:"context,attr"`
	Let     []let    `xml:"let"`
	Assert  []assert `xml:"assert"`
}
type assert struct {
	Test    string `xml:"test,attr"`
	Message string `xml:",innerxml"`
}
type ns struct {
	Uri string `xml:"uri,attr"`
}
type let struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func emitApply(t *Node, c *rule) {

	name := t.child[0].name[1:]
	r, _ := symtab[name]

	a := assert{
		Test:    r.value,
		Message: t.value,
	}

	c.Assert = append(c.Assert, a)

}

func emitContextBody(t []*Node, c *rule) {

	for _, n := range t {
		if n.nodeType == nodeVariable {
			emitVariable(n, &c.Let)
		} else {
			emitApply(n, c)
		}
	}
}

func emitContext(t *Node, p *pattern) {

	contextStr := ""
	if t.child[0].nodeType == nodeReference {
		name := t.child[0].name[1:]
		contextStr = symtab[name].value
	} else {
		contextStr = t.child[0].value
	}

	r := rule{
		Context: contextStr,
	}
	emitContextBody(t.child[1].child, &r)
	p.Rule = append(p.Rule, r)

}

func emitVariable(t *Node, a *[]let) {

	value := symtab[t.name].value
	v := let{
		Name:  t.name,
		Value: value,
	}
	*a = append((*a), v)

}
func generateCode(t []*Node) {
	s := &schema{}
	p := pattern{Name: "main"}

	for _, n := range t {
		switch n.nodeType {
		case nodeVariable:
			emitVariable(n, &s.Let)
		case nodeContext:
			emitContext(n, &p)
		}

	}

	s.Pattern = append(s.Pattern, p)
	output, err := xml.MarshalIndent(s, " ", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(html.UnescapeString(string(output)))

}
