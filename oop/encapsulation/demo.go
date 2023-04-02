package main

import "fmt"

// 封装

type Animal struct {
	name string
}

func NewAnimal() *Animal {
	return &Animal{}
}

func (p *Animal) SetName(name string) {
	p.name = name
}

func (p *Animal) GetName() string {
	return p.name
}

func main() {
	a := NewAnimal()
	a.SetName("动物")
	fmt.Println(a.GetName())
}