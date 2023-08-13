package main

import (
	"fmt"
	"unsafe"
)

// 利用空结构体实现 set 集合
func main()  {
	set := map[string]struct{}{
		"pm": {},
		"fe": {},
		"rd": {},
	}

	if v, ok := set["rd"]; ok {
		fmt.Println("exist")
		fmt.Println("size:", unsafe.Sizeof(v))
	} else {
		fmt.Println("not exist")
	}
}