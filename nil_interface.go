package main

import (
	"fmt"
)

type S struct{}

func (s S) F() {}

type IF interface {
	F()
}

func InitType() S {
	var s S
	return s
}

func InitPointer() *S {
	var s *S
	return s
}
func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}

func InitIfacePointer2() IF {
	var i IF
	return i
}

func main() {
	fmt.Println(InitType() == nil)
	fmt.Println(InitPointer() == nil)
	fmt.Println(InitEfaceType() == nil)
	fmt.Println(InitEfacePointer() == nil)
	fmt.Println(InitIfaceType() == nil)
	fmt.Println(InitIfacePointer() == nil)
	fmt.Println(InitIfacePointer2() == nil)
}
