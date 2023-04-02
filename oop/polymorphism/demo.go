package main

import "fmt"

// 多态 - 不同类实现接口呈现多态

type IAnimal interface {
	Say()
}

type Dog struct {}

func (d *Dog) Say() {
	fmt.Println("汪汪叫")
}

type Cat struct {}

func (c *Cat) Say() {
	fmt.Println("喵喵叫")
}

func main()  {
	var dog, cat IAnimal = &Dog{}, &Cat{}
	dog.Say()
	cat.Say()
}