package main

import (
	"fmt"
	"math/big"
)

// WeChatPayStrategy 具体策略，实现微信支付
type WeChatPayStrategy struct{}

func (w *WeChatPayStrategy) Pay(amount *big.Float) {

	fmt.Printf("使用微信支付：%s元\n", amount.Text('f', 2))
}
