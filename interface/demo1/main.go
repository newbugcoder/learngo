package main

import "github.com/newbugcoder/learngo/interface/demo1/lib"

func main()  {
	way := "we"
	//way := "ali"
	var payService lib.IPayService = lib.NewIPay(way)

	payService.Pay()
}