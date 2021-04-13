package compiler

import "regexp"

var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
