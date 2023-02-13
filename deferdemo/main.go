package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//demo0()
	//demo1()
	//demo2()
	//demo3_1()
	//demo3_2()
	//demo3_3()
	//demo4_1()
	//fmt.Println(demo4_2())
	//fmt.Println(demo4_3())
	//demo5_1()
	demo5_2()
}

// defer 常规用法
func demo0() {
	fileName := "./test.txt"
	f, _ := os.OpenFile(fileName, os.O_RDONLY, 0)
	defer f.Close()

	contents, _ := ioutil.ReadAll(f)
	fmt.Println(string(contents))
}

// 1. 展示多个 defer 的调用顺序："先进后出"
func demo1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("defer:", i)
	}
}

// 2. defer 作用域为当前函数，在当前函数最后执行。不同函数下拥有不同的 defer 栈。
func demo2() {
	func() {
		defer fmt.Println(1)
		defer fmt.Println(2)
	}()

	fmt.Println("=== 新生代农民工啊 ===")

	func() {
		defer fmt.Println("a")
		defer fmt.Println("b")
	}()
}

// 3. defer 传入函数的形参在声明时就已确定（预计算参数），而不是在执行时确定；

// defer 声明时，已经确认了形参n的值；后续 num 变量无论如何改变都不影响 defer 的输出结果。
func demo3_1() {
	num := 0
	defer func(n int) {
		fmt.Println("defer:", n)
	}(num)
	// 等同 defer fmt.Println("defer:", num)

	for i := 0; i < 10; i++ {
		num++
	}

	fmt.Println(num)

	//10
	//defer: 0
}

// defer 声明时，已经确认了形参p指针的指向地址，指向num变量；后续 num 变量发生改变。所以在 defer 执行时，输出的是p指针指向的num变量的当前值。
func demo3_2() {
	num := 0
	p := &num
	defer func(p *int) {
		fmt.Println("defer:", *p)
	}(p)

	for i := 0; i < 10; i++ {
		num++
	}

	fmt.Println(*p)

	//10
	//defer: 10
}

// defer 打印的变量并没有通过函数参数传入，而是在defer执行时，直接获取的"全局变量"。
func demo3_3() {
	num := 0
	defer func() {
		fmt.Println("defer:", num)
	}()

	for i := 0; i < 10; i++ {
		num++
	}

	fmt.Println(num)

	//10
	//defer: 10
}

// 4. return 与 defer 执行顺序，return 先，defer 后；

// 通过 demo4_1 可以发现 defer 在 return 之后执行；
func demo4_1() (int, error) {
	defer fmt.Println("defer")
	return fmt.Println("return")
}

// 围绕着 return、defer 的执行顺序和函数返回值，又会产生许多复杂的场景。

// demo4_2 函数使用命名返回值
// 1. 变量 num 作为返回值，初始值为0；
// 2. 随后变量 num 被赋值为 10；
// 3. return 时，变量 num 作为返回值被重新赋值为 2；
// 4. defer 在 return 后执行，拿到变量 num 进行修改，值为7；
// 5. 变量 num 作为返回值，最终函数返回结果为7；
func demo4_2() (num int) {
	num = 10
	defer func() {
		num += 5
	}()

	return 2
	// 7
}

//demo4_3 函数使用匿名返回值
//1. 此时返回值变量并未创建；
//2. 创建变量 num，赋值为 10；
//3. return 时创建函数返回值变量，并赋值为2；这个返回值变量你可以把它看成匿名变量，或者是变量 a、b、c、d……，但它就不是变量num；
//4. defer 时，无论怎样修改变量 num，都与函数返回值无关；
//5. 所以，最终的函数返回结果为2；
func demo4_3() int {
	num := 10
	defer func() {
		num += 5
	}()

	return 2
	// 2
}

// 5. 发生 panic 时，已声明的 defer 会出栈执行，再 panic；
func demo5_1() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	panic("异常") // 触发defer出栈执行

	defer fmt.Println(4) // 得不到执行
}

// 所以利用这个特性，在 defer 中可以 recover panic；
func demo5_2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("异常") // 触发defer出栈执行

	// ...
}
