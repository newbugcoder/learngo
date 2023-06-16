package main

import (
	"fmt"
	"strconv"
)

func main()  {
	demo1()
	demo2()
	demo3()
}

func demo1()  {
	var i interface{} = 2
	var param int = i.(int)
	square(param) // output:4

	var f float64 = 3
	var param2 int = int(f)
	square(param2) // output:9
}

func square(i int)  {
	fmt.Println(i * i)
}

// 类型转换
func demo2()  {
	var s string = "123"
	var num int
	num, _ = strconv.Atoi(s)
	fmt.Println(num)
}

// 断言
func demo3()  {
	var i interface{} = 2
	num1, ok := i.(int)
	fmt.Println(num1)

	var i2 interface{} = "新生代农民工啊"
	num2, ok := i2.(int)
	if ok {
		fmt.Println(num2, "断言成功")
	} else {
		fmt.Println(num2, "断言失败，它不是int类型")
	}

	var i3 interface{}
	i = "新生代农民工啊"
	output(i3)
	i3 = 123
	output(i3)
}

func output(i interface{})  {
	switch i.(type) {
	case string:
		fmt.Println("我现在是字符串类型")
	case int:
		fmt.Println("我现在是int类型")
	}
}