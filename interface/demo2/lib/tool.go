package lib

import "fmt"

type IResponse interface {
	MakeText() string
}

func Output(resp IResponse)  {
	text := resp.MakeText()
	fmt.Println("将此内容输出:", text)
}
