package test_github_private_package

import "fmt"

var Abc abc

type abc struct {
}

func (abc abc) Echo() {
	fmt.Print("abc")
}