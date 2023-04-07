package main

import "github.com/newbugcoder/learngo/interface/demo2/lib"

type xmlResp struct {}

func (xml *xmlResp) MakeText() string {
	return "组装成xml格式的数据"
}

type jsonResp struct {}

func (json *jsonResp) MakeText() string {
	return "组装成json格式的数据"
}

type sliceStr []string

func (sl sliceStr) MakeText() string {
	return "组装成切片格式的数据"
}

func main()  {
  lib.Output(&xmlResp{})
  lib.Output(&jsonResp{})
  sl := sliceStr{}
  lib.Output(sl)
}