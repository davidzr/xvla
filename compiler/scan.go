package compiler

import (
	"regexp"
)

var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString

type StateType int

var position int
var source string

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

func SetSource(s string) {
	source = s
}
func NextChar() (result *string) {
	if position < len(source) {
		char := string(source[position])
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
			} else if char == "/" {
				state = ENTERINGCOMMENT
				save = false
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
		case ENTERINGCOMMENT:
			save = false
			if char == "*" {
				state = INCOMMENT
			}
		case INCOMMENT:
			save = false
			if char == "*" {
				state = EXITINGCOMMENT
			}
		case EXITINGCOMMENT:
			save = false
			if char == "/" {
				state = START
			} else if char != "*" {
				state = INCOMMENT
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
	return currentToken, TokenString
}
