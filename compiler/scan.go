package compiler

import (
	"fmt"
	"regexp"
)

var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

type StateType int

var position int

const (
	START = iota
	INIDENTIFIER
	INREFERENCE
	INSTRING
	DONE
)

func NextChar() (result *string) {
	text := "ns hola = \"una cosa.com\"; context($13f1){var i = \"78\";var 7q = \"5\";apply $wbg{return \"asd\"}}rule g{assert \"x\"}var 7a = \"3\";"
	if position < len(text) {
		char := string(text[position])
		position++
		return &char
	} else {
		char := ""
		return &char
	}
}
func UnGetNextChar() {
	if position > 0 {
		position--
	}
}
func NextToken() (TokenType, string) {
	save := false
	var TokenString string
	state := START
	var currentToken TokenType
	for state != DONE {

		char := *NextChar()
		save = true
		switch state {
		case START:

			if isAlphaNumeric(char) {
				state = INIDENTIFIER
			} else if char == "$" {
				state = INREFERENCE
			} else if char == "\"" {
				state = INSTRING
			} else if char == "\n" || char == " " {
				save = false
			} else {
				state = DONE
				switch char {
				case "":
					currentToken = EOF
					save = false
				case "(":
					currentToken = LPARENT
				case ")":
					currentToken = RPARENT
				case "{":
					currentToken = LBRACKET
				case "}":
					currentToken = RBRACKET
				case "=":
					currentToken = EQUAL
				case ";":
					currentToken = SEMICOLON
				default:
					currentToken = ERROR
				}
			}
		case INIDENTIFIER:
			if !isAlphaNumeric(char) {
				UnGetNextChar()
				save = false
				state = DONE
				currentToken = IDENTIFIER
			}
		case INREFERENCE:
			if !isAlphaNumeric(char) {
				UnGetNextChar()
				save = false
				state = DONE
				currentToken = REFERENCE
			}
		case INSTRING:
			if char == "\"" {
				state = DONE
				currentToken = STRING
			}
		default:
			state = DONE
			currentToken = ERROR
		}

		if save {
			TokenString += char
		}
		if state == DONE {
			if currentToken == IDENTIFIER {
				t, ok := ReservedWords[TokenString]
				if ok {
					currentToken = t
				}
			}
		}
	}
	fmt.Println("--", TokenString)
	return currentToken, TokenString
}
