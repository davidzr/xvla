package main

import (
	"io/ioutil"

	"github.com/davidzr/xvla/compiler"
)

var source string

func main() {
	source, err := ioutil.ReadFile("/mnt/d/Desarrollo/www/dianfe/ejemplo.vvxxd")
	if err != nil {
		panic(err)
	}
	t := compiler.Parse(string(source))
	compiler.Analyze(t.Sibling)
}
