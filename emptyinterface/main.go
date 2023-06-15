package main

import "fmt"

func main()  {
	//demo1()
	//demo2()
	demo3()
}

func demo1()  {
	var i interface{}

	i = []int{1,2,3}
	fmt.Println(i)
	test(i)

	i = "abc"
	fmt.Println(i)
	test(i)
}

func test(foo interface{})  {
	switch foo.(type) {
	case string:
		str := foo.(string)
		fmt.Printf("%T\n", str) // string
	case []int:
		sl := foo.([]int)
		fmt.Printf("%T\n", sl) // []int
	}
}

func demo2()  {
	var i interface{}
	var p *interface{} = &i
	fmt.Println(p)
	num := 11
	var numsP *int = &num
	fmt.Println(*numsP)
	//p = numsP // 报错，不能把 *int 赋值给 *interface
}

func demo3()  {
	var sl []interface{}
	var sl2 = []int{1, 2, 3}
	//sl = sl2 // 报错，不能直接赋值

	// 需要将每个数据项逐个赋值
	for _, num := range sl2 {
		sl = append(sl, num)
	}
	fmt.Println(sl) // [1 2 3]
}