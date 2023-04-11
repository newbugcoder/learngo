package main

import "github.com/newbugcoder/learngo/interface/qa/lib"

type Dog struct {}

func (d *Dog) Get() string {
	return "dog"
}

func fDog() string {
	return "fDog"
}

func main()  {
	lib.Output(&Dog{})

	lib.Output2(func() string {
		return "匿名函数dog"
	})

	lib.Output2(fDog)
}

//dog
//匿名函数dog
//fDog