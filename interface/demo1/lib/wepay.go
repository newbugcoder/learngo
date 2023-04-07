package lib

import "fmt"

type WePay struct {}

func NewWePay() IPayService {
	return &WePay{}
}

func (w *WePay) Pay() bool {
	fmt.Println("wepay...")
	// 组装参数，调用接口
	return true
}
