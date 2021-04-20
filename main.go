package main

import (
	"io/ioutil"

	"github.com/davidzr/xvla/compiler"
)

var source string

func main() {
	source, err := ioutil.ReadFile("/mnt/d/Desarrollo/www/dianfe/test2.txt")
	if err != nil {
		panic(err)
	}
	compiler.Run(string(source))
}
