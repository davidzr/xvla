package compiler

import (
	"fmt"
	"regexp"
)

var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString

type StateType int

var position int

const (
	START = iota
	INLITERAL
	INIDENTIFIER
	INSTRING
	DONE
)

func NextChar() (result *string) {
	text := "var ns = \"hola\"; rule hola{return \"prueba\"}"
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
	position--
}
func NextToken() TokenType {
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
				state = INLITERAL
			} else if char == "$" {
				state = INIDENTIFIER
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
		case INLITERAL:
			if !isAlphaNumeric(char) {
				UnGetNextChar()
				save = false
				state = DONE
				currentToken = LITERAL
			}
		case INIDENTIFIER:
			if !isAlphaNumeric(char) {
				UnGetNextChar()
				save = false
				state = DONE
				currentToken = IDENTIFIER
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
			if currentToken == LITERAL {
				t, ok := ReservedWords[TokenString]
				if ok {
					currentToken = t
				}
			}
		}
	}
	fmt.Println(TokenString)
	return currentToken
}
