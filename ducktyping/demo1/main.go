package main

import (
	"github.com/newbugcoder/learngo/ducktyping/demo1/pkga"
	"github.com/newbugcoder/learngo/ducktyping/demo1/pkgb"
)

func main()  {
	classB := pkgb.NewClassB()
	pkga.Test(classB)
}