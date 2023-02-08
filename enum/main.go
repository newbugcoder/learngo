package main

import (
	"fmt"
	"github.com/newbugcoder/learngo/enum/Weekday"
)

func main() {
	var sunday = Weekday.Sunday
	fmt.Println(sunday)          // Print : Sunday
	fmt.Println(sunday.String()) // Print : Sunday
	fmt.Println(sunday.Index())  // Print : 1

	fmt.Println(Weekday.Values())
	// Print : [Sunday Monday Tuesday Wednesday Thursday Friday Saturday]
	fmt.Println(Weekday.ExistOf("abc"))    // Print : false
	fmt.Println(Weekday.ExistOf("Monday")) // Print : true
}