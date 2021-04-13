package compiler

import (
	"testing"
)

func TestVariable(t *testing.T) {
	s := &scanner{
		source: `var i = 123`,
	}
	s.nextToken()

	if s.token != VARIABLE {
		t.Fatal("nextToken() must be variable.")
	}

}
