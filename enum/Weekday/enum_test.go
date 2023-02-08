package Weekday

import (
	"fmt"
	"testing"
)

func TestExistOf(t *testing.T) {
	// 判断
	str := "abc"
	fmt.Printf("%#v\n", ExistOf(str))
}