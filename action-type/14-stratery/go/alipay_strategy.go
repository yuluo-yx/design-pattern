package main

import (
	"fmt"
	"math/big"
)

// AliPayStrategy 具体策略，实现支付宝支付
type AliPayStrategy struct{}

func (a *AliPayStrategy) Pay(amount *big.Float) {

	fmt.Printf("使用支付宝支付：%s元\n", amount.Text('f', 2))
}
