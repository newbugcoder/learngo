package main

import (
	"fmt"
	"unsafe"
)
// 空结构体作为方法接收器
type Client struct {}

func NewClient() Client {
	return Client{}
}

func (c Client) Send() {
	fmt.Println("send")
}

func main()  {
	c := NewClient()
	fmt.Println("size:", unsafe.Sizeof(c))
	c.Send()
}