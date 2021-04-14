package compiler

import (
	"testing"
)

func TestRule(t *testing.T) {
	s := &scanner{
		source: `rule test{assert "$1 = 1"}`,
	}
	s.nextToken()

	if s.token != RULE {
		t.Fatal("nextToken() must be rule.")
	}

}
func TestVariable(t *testing.T) {
	s := &scanner{
		source: `var i = 123`,
	}
	s.nextToken()

	if s.token != VARIABLE {
		t.Fatal("nextToken() must be variable.")
	}

}
