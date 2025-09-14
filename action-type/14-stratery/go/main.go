package main

import (
	"math/big"
)

func main() {

	context := &PayContext{}

	// 用户选择支付宝
	context.SetStrategy(&AliPayStrategy{})
	context.Pay(big.NewFloat(100.00))

	// 用户选择微信
	context.SetStrategy(&WeChatPayStrategy{})
	context.Pay(big.NewFloat(200.00))
}
