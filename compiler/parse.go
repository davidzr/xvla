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

func (p *parser) variableStmt() *Node {
	t := NewNode(nodeVariable, p.s.line)
	p.match(VARIABLE)
	t.name = p.s.tokenString
	p.match(IDENTIFIER)
	p.match(EQUAL)
	t.child = append(t.child, p.stringStmt())
	p.match(SEMICOLON)

	return t
}

func (p *parser) stringStmt() *Node {
	t := NewNode(nodeLiteral, p.s.line)
	t.value = p.s.tokenString
	p.match(STRING)
	return t
}

func (p *parser) referenceStmt() *Node {
	t := NewNode(nodeReference, p.s.line)
	t.name = p.s.tokenString
	p.match(REFERENCE)
	return t
}

func (p *parser) namespaceStmt() *Node {
	t := NewNode(nodeNamespace, p.s.line)
	p.match(NAMESPACE)
	t.name = p.s.tokenString
	p.match(IDENTIFIER)
	p.match(EQUAL)
	t.child = append(t.child, p.stringStmt())
	p.match(SEMICOLON)

	return t
}

func (p *parser) returnStmt() *Node {
	t := NewNode(nodeReturn, p.s.line)
	p.match(RETURN)
	t.child = append(t.child, p.stringStmt())
	return t
}

func (p *parser) applyStmt() *Node {
	t := NewNode(nodeApply, p.s.line)
	p.match(APPLY)
	t.child = append(t.child, p.referenceStmt())
	p.match(LBRACKET)
	t.child = append(t.child, p.returnStmt())
	p.match(RBRACKET)
	return t
}

func (p *parser) contextBody() *Node {
	t := NewNode(nodeContextBody, p.s.line)
	for p.s.token == VARIABLE || p.s.token == APPLY {
		switch p.s.token {
		case VARIABLE:
			t.child = append(t.child, p.variableStmt())
		case APPLY:
			t.child = append(t.child, p.applyStmt())
		}
	}
	return t
}

func (p *parser) contextStmt() *Node {
	t := NewNode(nodeContext, p.s.line)
	p.match(CONTEXT)
	p.match(LPARENT)
	if p.s.token == REFERENCE {
		t.child = append(t.child, p.referenceStmt())
	} else if p.s.token == STRING {
		t.child = append(t.child, p.stringStmt())
	}
	p.match(RPARENT)
	p.match(LBRACKET)
	t.child = append(t.child, p.contextBody())
	p.match(RBRACKET)
	return t
}
func (p *parser) assertStmt() *Node {
	t := NewNode(nodeAssert, p.s.line)
	p.match(ASSERT)
	t.child = append(t.child, p.stringStmt())
	return t
}

func (p *parser) ruleBody() *Node {
	t := NewNode(nodeRuleBody, p.s.line)
	for p.s.token == VARIABLE {
		t.child = append(t.child, p.variableStmt())
	}
	t.child = append(t.child, p.assertStmt())

	return t
}

func (p *parser) ruleStmt() *Node {
	t := NewNode(nodeRule, p.s.line)
	p.match(RULE)
	t.name = p.s.tokenString
	p.match(IDENTIFIER)
	p.match(LBRACKET)
	t.child = append(t.child, p.ruleBody())
	p.match(RBRACKET)
	return t
}
func (p *parser) statement() *Node {
	var t *Node
	if p.s.token == CONTEXT {
		t = p.contextStmt()
	} else if p.s.token == RULE {
		t = p.ruleStmt()
	} else if p.s.token == VARIABLE {
		t = p.variableStmt()
	} else if p.s.token == NAMESPACE {
		t = p.namespaceStmt()
	}
	return t
}
func (p *parser) program() *Node {
	t := p.statement()
	program := NewNode(nodeProgram, p.s.line)
	program.sibling = append(program.sibling, t)
	for p.s.token == CONTEXT || p.s.token == VARIABLE || p.s.token == NAMESPACE || p.s.token == RULE {
		n := p.statement()
		program.sibling = append(program.sibling, n)
	}
	return program
}

func parse(source string) *Node {
	s := &scanner{source: source}
	p := &parser{s: s}
	p.s.nextToken()
	t := p.program()
	return t
}
