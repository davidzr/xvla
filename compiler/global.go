package compiler

type TokenType int

const (
	CONTEXT = iota
	RULE
	APPLY
	RETURN
	NAMESPACE
	ASSERT
	VARIABLE
	IDENTIFIER
	EQUAL
	LPARENT
	RPARENT
	RBRACKET
	LBRACKET
	SEMICOLON
	STRING
	LITERAL
	EOF
	ERROR
)

var ReservedWords = map[string]TokenType{
	"rule":    RULE,
	"appply":  APPLY,
	"return":  RETURN,
	"ns":      NAMESPACE,
	"assert":  ASSERT,
	"var":     VARIABLE,
	"context": CONTEXT,
}