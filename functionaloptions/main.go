package main

import (
	co "github.com/newbugcoder/learngo/functionaloptions/configobject"
	"github.com/newbugcoder/learngo/functionaloptions/lib1"
	"github.com/newbugcoder/learngo/functionaloptions/lib2"
)

// 使用

func main()  {
	// 配置对象方案
	co.NewServer("127.0.0.1", 8080, nil)
	co.NewServer("127.0.0.1", 8080, &co.Config{})
	co.NewServer("127.0.0.1", 8080, &co.Config{
		Protocol: "UDP",
	})
	// 功能选项方案-基于闭包函数
	lib1.NewServer("127.0.0.1", 8080)
	lib1.NewServer("127.0.0.1", 8080, lib1.WithProtocol("UDP"))
	// 功能选项方案-基于接口
	lib2.NewServer("127.0.0.1", 8080)
	lib2.NewServer("127.0.0.1", 8080, lib2.WithProtocol("UDP"))
}