package main

import (
	"fmt"
	"unsafe"
)

// 空结构体不占用空间

type Empty struct {

}

//func main()  {
//	a := Empty{}
//	b := struct{}{}
//
//	fmt.Printf("addr: %p, size: %d\n", &a, unsafe.Sizeof(a))
//	fmt.Printf("addr: %p, size: %d\n", &b, unsafe.Sizeof(b))
//}

type MyStruct struct {
	Foo Empty
	Bar struct{}
}

func main()  {
	a := Empty{}
	my := MyStruct{}
	fmt.Printf("addr: %p, size: %d\n", &a, unsafe.Sizeof(a))
	fmt.Printf("addr: %p, size: %d\n", &my, unsafe.Sizeof(my))
	fmt.Printf("addr: %p, size: %d\n", &my.Foo, unsafe.Sizeof(my.Foo))
}