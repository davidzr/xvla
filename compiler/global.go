package compiler

type TokenType int
type StateType int

const (
	CONTEXT TokenType = iota
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
	START StateType = iota
	INIDENTIFIER
	INREFERENCE
	INSTRING
	ENTERINGCOMMENT
	EXITINGCOMMENT
	INCOMMENT
	DONE
)

var keywords = map[string]TokenType{
	"rule":    RULE,
	"apply":   APPLY,
	"return":  RETURN,
	"ns":      NAMESPACE,
	"assert":  ASSERT,
	"var":     VARIABLE,
	"context": CONTEXT,
}
