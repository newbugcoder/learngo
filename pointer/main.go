package main

import "fmt"

func main()  {
	demo1()
	//demo2()
}

func demo1()  {
	var str string = "code"
	fmt.Println("str:", str)
	fmt.Println("&str", &str)

	var ptr *string = &str // 初始化指针
	fmt.Println("ptr:", ptr)
	fmt.Println("*ptr", *ptr)
}

func demo2()  {
	var str string = "code"
	fmt.Println("str:", str)

	var ptr *string = &str
	fmt.Println("*ptr:", *ptr)

	fmt.Println("=== 修改后 ===")
	*ptr = "新生代农民工啊"
	fmt.Println("str:", str)
	fmt.Println("*ptr:", *ptr)
}