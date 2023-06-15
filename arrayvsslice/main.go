package main

import "fmt"

func main()  {
	// array
	var av = [5]int{1,5,2,3,7}
	var bv = [...]int{1,5,2,3,7}
	// slice
	var cv = []int{1,5,2,3,7}

	//判断变量类型
	fmt.Printf("%T\n", av)
	fmt.Printf("%T\n", bv)
	fmt.Printf("%T\n", cv)
}
