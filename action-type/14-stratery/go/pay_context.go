package main

import "math/big"

// PayContext 上下文，持有当前的支付策略
// 可以动态切换
type PayContext struct {
	strategy PayStrategy
}

// SetStrategy 设置当前支付策略
func (c *PayContext) SetStrategy(strategy PayStrategy) {

	c.strategy = strategy
}

// Pay 执行支付动作
func (c *PayContext) Pay(amount *big.Float) {

	if c.strategy == nil {
		panic("未设置支付策略")
	}

	c.strategy.Pay(amount)
}
