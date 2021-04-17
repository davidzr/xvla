package compiler

import (
	"encoding/xml"
	"fmt"
	"os"
)

type schema struct {
	Ns      []ns      `xml:"ns"`
	Pattern []pattern `xml:"pattern"`
	Let     []let     `xml:"let"`
}

type pattern struct {
	Rule []rule `xml:"rule"`
	Name string `xml:"name,attr"`
}
type rule struct {
	Context string   `xml:"context,attr"`
	Assert  []assert `xml:"assert"`
}
type assert struct {
	Test string `xml:"test,attr"`
}
type ns struct {
	Uri string `xml:"uri,attr"`
}
type let struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func emitContext(t *Node, p *pattern) {

	contextStr := ""
	if t.child[0].nodeType == nodeReference {
		name := t.child[0].name[1:]
		contextStr = symtab[name].value
	} else {
		contextStr = t.child[0].value
	}
	fmt.Println(contextStr)
	r := rule{
		Context: contextStr,
	}
	p.Rule = append(p.Rule, r)

}

func generateCode(t []*Node) {
	s := &schema{}
	p := pattern{Name: "main"}

	for _, n := range t {
		switch n.nodeType {
		case nodeContext:
			emitContext(n, &p)
		}
	}

	s.Pattern = append(s.Pattern, p)
	output, err := xml.Marshal(s)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
