package main

import (
	"fmt"
	"time"
)

// 空结构体作为 channel 信号
func main()  {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		ch <- struct{}{}
	}()

	fmt.Println("waiting...")
	<-ch
	fmt.Println("end")
}
