package main

// 继承

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Run() {
	fmt.Println(a.Name,"会跑")
}

type Dog struct {
	Animal
}

////多态 - 子类重写父类同名方法呈现多态
//func (d *Dog) Run() {
//	fmt.Println(d.Name,"撒欢跑")
//}

type Cat struct {
	Animal
}

////多态
//func (c *Cat) Run() {
//	fmt.Println(c.Name,"走猫步")
//}

func main()  {
	dog := Dog{Animal{Name: "大黄"}}
	dog.Run()

	cat := Cat{Animal{Name: "喵喵"}}
	cat.Run()
}