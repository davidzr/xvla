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
	REFERENCE
	EOF
	ERROR
)

const (
	START = iota
	INIDENTIFIER
	INREFERENCE
	INSTRING
	ENTERINGCOMMENT
	EXITINGCOMMENT
	INCOMMENT
	DONE
)

var ReservedWords = map[string]TokenType{
	"rule":    RULE,
	"apply":   APPLY,
	"return":  RETURN,
	"ns":      NAMESPACE,
	"assert":  ASSERT,
	"var":     VARIABLE,
	"context": CONTEXT,
}
