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
	match(IDENTIFIER)
	match(EQUAL)
	match(STRING)
	fmt.Println("variable", tokenString)
	match(SEMICOLON)
}

func return_stmt() {
	match(RETURN)
	match(STRING)
	fmt.Println("return", tokenString)
}

func apply_stmt() {
	match(APPLY)
	match(REFERENCE)
	match(LBRACKET)
	return_stmt()
	match(RBRACKET)
}

func context_body() {
	for token == VARIABLE || token == APPLY {
		if token == VARIABLE {
			variable_stmt()
		} else {
			apply_stmt()
		}
	}
}

func context_stmt() {
	match(CONTEXT)
	match(LPARENT)
	if token == REFERENCE {
		fmt.Println("LLego ID", tokenString)
		match(REFERENCE)
	} else if token == STRING {
		fmt.Println("LLego STRING")
		match(STRING)
	}
	match(RPARENT)
	match(LBRACKET)
	context_body()
	match(RBRACKET)
}
func assert_stmt() {
	match(ASSERT)
	match(STRING)
	fmt.Println("assert", tokenString)
}

func rule_body() {
	assert_stmt()
}

func rule_stmt() {
	match(RULE)
	match(IDENTIFIER)
	match(LBRACKET)
	rule_body()
	match(RBRACKET)
}

func program() {
	for token == CONTEXT || token == VARIABLE || token == RULE {
		if token == CONTEXT {
			context_stmt()
		} else if token == RULE {
			rule_stmt()
		} else if token == VARIABLE {
			variable_stmt()
		}
	}
}

func Parse() {
	token, tokenString = NextToken()
	program()
}
