package pkgb

import "fmt"

type ClassB struct {}

func NewClassB() *ClassB {
	return &ClassB{}
}

func (c ClassB) DoSomething() {
	fmt.Println("ClassB")
}