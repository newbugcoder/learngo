package main

import "fmt"

type Typer interface {
	Who() string
}

type Person struct {}

func (p *Person) Who() string {
	return "我是Person结构体"
}

type strsl []string

func (s strsl) Who() string {
	return "我是字符串切片"
}

type in int

func (i in) Who() string {
	return "我是整形"
}

func WoAmI(t Typer)  {
	fmt.Println(t.Who())
}

func main()  {
	p := &Person{}
	WoAmI(p) // 我是Person结构体
	var sl strsl = []string{"a"}
	WoAmI(sl) // 我是字符串切片
	var i in = 1
	WoAmI(i) // 我是整形
}
