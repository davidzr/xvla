package compiler

import (
	"fmt"
)

var token TokenType
var tokenString string

func match(expected TokenType) {
	token, tokenString = NextToken()
	if token != expected {
		fmt.Println(expected, token)
		panic("unexpected token " + tokenString)
	}
}

func variable_stmt() {
	match(LITERAL)
	match(EQUAL)
	match(STRING)
	fmt.Println("variable", tokenString)
	match(SEMICOLON)

	token, tokenString = NextToken()
}

func return_stmt() {
	match(RETURN)
	match(STRING)
	fmt.Println("return", tokenString)
}

func apply_stmt() {

	match(IDENTIFIER)
	match(LBRACKET)
	return_stmt()
	match(RBRACKET)
	token, tokenString = NextToken()

}

func context_body() {
	token, tokenString = NextToken()
	for token == VARIABLE || token == APPLY {
		fmt.Println("The end body", tokenString)
		if token == VARIABLE {
			variable_stmt()
		} else {
			apply_stmt()
		}
	}

}

func context_stmt() {
	match(LPARENT)
	token, tokenString = NextToken()
	if token == IDENTIFIER {

		fmt.Println("LLego ID", tokenString)
	} else if token == STRING {
		fmt.Println("LLego STRING")
	} else {
		panic("Unexpected token " + tokenString)
	}

	match(RPARENT)
	match(LBRACKET)
	context_body()
	token, tokenString = NextToken()
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
	match(LITERAL)
	match(LBRACKET)
	rule_body()
	match(RBRACKET)
	token, _ = NextToken()
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
