package pkga

type MyInterface interface {
	DoSomething()
}

func Test(m MyInterface) {
	m.DoSomething()
}