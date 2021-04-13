package compiler

type scanner struct {
	position            int
	line                int
	source, tokenString string
	token               TokenType
}

func (s *scanner) nextChar() (result *string) {
	if s.position < len(s.source) {
		char := string(s.source[s.position])
		s.position++
		if char == "\n" {
			s.line++
		}
		return &char
	} else {
		char := ""
		return &char
	}
}
func (s *scanner) unGetNextChar() {
	if s.position > 0 {
		s.position--
	}
}
func (s *scanner) nextToken() {
	save := false
	state := START
	s.tokenString = ""
	for state != DONE {

		char := *s.nextChar()
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
			} else if char == "\n" {
				save = false
			} else if char == " " {
				save = false
			} else {
				state = DONE
				switch char {
				case "":
					s.token = EOF
					save = false
				case "(":
					s.token = LPARENT
				case ")":
					s.token = RPARENT
				case "{":
					s.token = LBRACKET
				case "}":
					s.token = RBRACKET
				case "=":
					s.token = EQUAL
				case ";":
					s.token = SEMICOLON
				default:
					s.token = ERROR
				}
			}
		case INIDENTIFIER:
			if !isAlphaNumeric(char) {
				s.unGetNextChar()
				save = false
				state = DONE
				s.token = IDENTIFIER
			}
		case INREFERENCE:
			if !isAlphaNumeric(char) {
				s.unGetNextChar()
				save = false
				state = DONE
				s.token = REFERENCE
			}
		case INSTRING:
			if char == "\"" {
				state = DONE
				s.token = STRING
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
			s.token = ERROR
		}

		if save {
			s.tokenString += char
		}
		if state == DONE {
			if s.token == IDENTIFIER {
				t, ok := ReservedWords[s.tokenString]
				if ok {
					s.token = t
				}
			}
		}
	}

}
