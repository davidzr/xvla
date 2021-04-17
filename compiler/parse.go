package compiler

import (
	"strconv"
)

type parser struct {
	s *scanner
}

func (p *parser) match(expected TokenType) {
	if p.s.token != expected {
		panic("unexpected token " + p.s.tokenString + ", On line:" + strconv.Itoa(p.s.line))
	}
	p.s.nextToken()
}

func (p *parser) variable_stmt() *Node {
	t := NewNode(nodeVariable, p.s.line)
	p.match(VARIABLE)
	t.name = p.s.tokenString
	p.match(IDENTIFIER)
	p.match(EQUAL)
	t.child = append(t.child, p.string_stmt())
	p.match(SEMICOLON)

	return t
}

func (p *parser) string_stmt() *Node {
	t := NewNode(nodeLiteral, p.s.line)
	t.value = p.s.tokenString
	p.match(STRING)
	return t
}

func (p *parser) reference_stmt() *Node {
	t := NewNode(nodeReference, p.s.line)
	t.name = p.s.tokenString
	p.match(REFERENCE)
	return t
}

func (p *parser) namespace_stmt() *Node {
	t := NewNode(nodeNamespace, p.s.line)
	p.match(NAMESPACE)
	t.name = p.s.tokenString
	p.match(IDENTIFIER)
	p.match(EQUAL)
	t.child = append(t.child, p.string_stmt())
	p.match(SEMICOLON)

	return t
}

func (p *parser) return_stmt() *Node {
	t := NewNode(nodeReturn, p.s.line)
	p.match(RETURN)
	t.child = append(t.child, p.string_stmt())
	return t
}

func (p *parser) apply_stmt() *Node {
	t := NewNode(nodeApply, p.s.line)
	p.match(APPLY)
	t.child = append(t.child, p.reference_stmt())
	p.match(LBRACKET)
	t.child = append(t.child, p.return_stmt())
	p.match(RBRACKET)
	return t
}

func (p *parser) context_body() *Node {
	t := NewNode(nodeContextBody, p.s.line)
	for p.s.token == VARIABLE || p.s.token == APPLY {
		switch p.s.token {
		case VARIABLE:
			t.child = append(t.child, p.variable_stmt())
		case APPLY:
			t.child = append(t.child, p.apply_stmt())
		}
	}
	return t
}

func (p *parser) context_stmt() *Node {
	t := NewNode(nodeContext, p.s.line)
	p.match(CONTEXT)
	p.match(LPARENT)
	if p.s.token == REFERENCE {
		t.child = append(t.child, p.reference_stmt())
	} else if p.s.token == STRING {
		t.child = append(t.child, p.string_stmt())
	}
	p.match(RPARENT)
	p.match(LBRACKET)
	t.child = append(t.child, p.context_body())
	p.match(RBRACKET)
	return t
}
func (p *parser) assert_stmt() *Node {
	t := NewNode(nodeAssert, p.s.line)
	p.match(ASSERT)
	t.child = append(t.child, p.string_stmt())
	return t
}

func (p *parser) rule_body() *Node {
	t := NewNode(nodeRuleBody, p.s.line)
	for p.s.token == VARIABLE {
		t.child = append(t.child, p.variable_stmt())
	}
	t.child = append(t.child, p.assert_stmt())

	return t
}

func (p *parser) rule_stmt() *Node {
	t := NewNode(nodeRule, p.s.line)
	p.match(RULE)
	t.name = p.s.tokenString
	p.match(IDENTIFIER)
	p.match(LBRACKET)
	t.child = append(t.child, p.rule_body())
	p.match(RBRACKET)
	return t
}
func (p *parser) statement() *Node {
	var t *Node
	if p.s.token == CONTEXT {
		t = p.context_stmt()
	} else if p.s.token == RULE {
		t = p.rule_stmt()
	} else if p.s.token == VARIABLE {
		t = p.variable_stmt()
	} else if p.s.token == NAMESPACE {
		t = p.namespace_stmt()
	}
	return t
}
func (p *parser) program() *Node {
	t := p.statement()
	for p.s.token == CONTEXT || p.s.token == VARIABLE || p.s.token == NAMESPACE || p.s.token == RULE {
		n := p.statement()
		t.sibling = append(t.sibling, n)
	}
	return t
}

func parse(source string) *Node {
	s := &scanner{source: source}
	p := &parser{s: s}
	p.s.nextToken()
	t := p.program()
	return t
}
