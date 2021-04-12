package compiler

import (
	"fmt"
	"strconv"
)

var token TokenType
var tokenString string

func match(expected TokenType) {
	if token != expected {
		fmt.Println(expected, token)
		panic("unexpected token " + tokenString + ", On line:" + strconv.Itoa(line))
	}
	token, tokenString = NextToken()
}

func variable_stmt() *Node {
	t := NewNode(NodeVariable, line)
	match(VARIABLE)
	t.child = append(t.child, identifier_stmt())
	match(EQUAL)
	t.child = append(t.child, string_stmt())
	match(SEMICOLON)

	return t
}

func string_stmt() *Node {
	t := NewNode(NodeString, line)
	t.value = tokenString
	match(STRING)
	return t
}

func reference_stmt() *Node {
	t := NewNode(NodeReference, line)
	t.value = tokenString
	match(REFERENCE)
	return t
}
func identifier_stmt() *Node {
	t := NewNode(NodeIdentifier, line)
	t.value = tokenString
	match(IDENTIFIER)
	return t
}

func namespace_stmt() *Node {
	t := NewNode(NodeNamespace, line)
	match(NAMESPACE)
	t.child = append(t.child, identifier_stmt())
	match(EQUAL)
	t.child = append(t.child, string_stmt())
	match(SEMICOLON)

	return t
}

func return_stmt() *Node {
	t := NewNode(NodeReturn, line)
	match(RETURN)
	t.child = append(t.child, string_stmt())
	return t
}

func apply_stmt() *Node {
	t := NewNode(NodeApply, line)
	match(APPLY)
	t.child = append(t.child, reference_stmt())
	match(LBRACKET)
	t.child = append(t.child, return_stmt())
	match(RBRACKET)
	return t
}

func context_body() *Node {
	t := NewNode(NodeContextBody, line)
	for token == VARIABLE || token == APPLY {
		switch token {
		case VARIABLE:
			t.child = append(t.child, variable_stmt())
		case APPLY:
			t.child = append(t.child, apply_stmt())
		}
	}
	return t
}

func context_stmt() *Node {
	t := NewNode(NodeContext, line)
	match(CONTEXT)
	match(LPARENT)
	if token == REFERENCE {
		t.child = append(t.child, reference_stmt())
	} else if token == STRING {
		t.child = append(t.child, string_stmt())
	}
	match(RPARENT)
	match(LBRACKET)
	t.child = append(t.child, context_body())
	match(RBRACKET)
	return t
}
func assert_stmt() *Node {
	t := NewNode(NodeAssert, line)
	match(ASSERT)
	t.child = append(t.child, string_stmt())
	return t
}

func rule_body() *Node {
	t := NewNode(NodeRuleBody, line)
	for token == VARIABLE {
		t.child = append(t.child, variable_stmt())
	}
	t.child = append(t.child, assert_stmt())

	return t
}

func rule_stmt() *Node {
	t := NewNode(NodeRule, line)
	match(RULE)
	t.child = append(t.child, identifier_stmt())
	match(LBRACKET)
	t.child = append(t.child, rule_body())
	match(RBRACKET)
	return t
}
func statement() *Node {
	var t *Node
	if token == CONTEXT {
		t = context_stmt()
	} else if token == RULE {
		t = rule_stmt()
	} else if token == VARIABLE {
		t = variable_stmt()
	} else if token == NAMESPACE {
		t = namespace_stmt()
	}
	return t
}
func program() Node {
	t := statement()
	for token == CONTEXT || token == VARIABLE || token == NAMESPACE || token == RULE {
		n := statement()
		t.Sibling = append(t.Sibling, n)
	}
	return *t
}

func Parse() Node {
	token, tokenString = NextToken()
	t := program()
	return t
}
