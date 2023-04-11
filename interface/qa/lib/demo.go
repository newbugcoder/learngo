package lib

import "fmt"

type IAnimal interface {
	Get() string
}

func Output(i IAnimal)  {
	fmt.Println(i.Get())
}

type FAnimal func() string

func Output2(f FAnimal)  {
	fmt.Println(f())
}