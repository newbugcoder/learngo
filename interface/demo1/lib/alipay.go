package lib

import "fmt"

type AliPay struct {}

func NewAliPay() IPayService {
	return &AliPay{}
}

func (a *AliPay) Pay() bool {
	fmt.Println("alipay...")
	// 组装参数，调用接口
	return true
}
