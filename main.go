package main

import (
	"fmt"
	"io/ioutil"

	"github.com/davidzr/xvla/compiler"
)

var source string

func main() {
	source, err := ioutil.ReadFile("/mnt/d/Desarrollo/www/dianfe/ejemplo.vvxxd")
	if err != nil {
		panic(err)
	}
	compiler.SetSource(string(source))
	compiler.Parse()
	fmt.Println("DOne")
}
