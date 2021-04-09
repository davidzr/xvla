package compiler

import (
	"fmt"
)

var token TokenType
var tokenString string

func match(expected TokenType) {
	if token != expected {
		fmt.Println(expected, token)
		panic("unexpected token " + tokenString)
	}
	token, tokenString = NextToken()
}

func variable_stmt() {
	match(VARIABLE)
	identifier_stmt()
	match(EQUAL)
	string_stmt()
	match(SEMICOLON)
}

func string_stmt() {
	match(STRING)
}

func reference_stmt() {
	match(REFERENCE)
}
func identifier_stmt() {
	match(IDENTIFIER)
}

func namespace_stmt() {
	match(NAMESPACE)
	identifier_stmt()
	match(EQUAL)
	string_stmt()
	match(SEMICOLON)
}

func return_stmt() {
	match(RETURN)
	string_stmt()
}

func apply_stmt() {
	match(APPLY)
	reference_stmt()
	match(LBRACKET)
	return_stmt()
	match(RBRACKET)
}

func context_body() {
	for token == VARIABLE || token == APPLY {
		switch token {
		case VARIABLE:
			variable_stmt()
		case APPLY:
			apply_stmt()
		}
	}
}

func context_stmt() {
	match(CONTEXT)
	match(LPARENT)
	if token == REFERENCE {
		reference_stmt()
	} else if token == STRING {
		string_stmt()
	}
	match(RPARENT)
	match(LBRACKET)
	context_body()
	match(RBRACKET)
}
func assert_stmt() {
	match(ASSERT)
	string_stmt()
}

func rule_body() {
	for token == VARIABLE {
		variable_stmt()
	}
	assert_stmt()
}

func rule_stmt() {
	match(RULE)
	identifier_stmt()
	match(LBRACKET)
	rule_body()
	match(RBRACKET)
}

func program() {
	for token == CONTEXT || token == VARIABLE || token == NAMESPACE || token == RULE {
		if token == CONTEXT {
			context_stmt()
		} else if token == RULE {
			rule_stmt()
		} else if token == VARIABLE {
			variable_stmt()
		} else if token == NAMESPACE {
			namespace_stmt()
		}
	}
}

func Parse() {
	token, tokenString = NextToken()
	program()
}
