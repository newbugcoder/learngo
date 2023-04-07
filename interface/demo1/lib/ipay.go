package lib

type IPayService interface {
	Pay() bool
}

func NewIPay(way string) IPayService {
	if way == "we" {
		return NewWePay()
	} else {
		return NewAliPay()
	}
}