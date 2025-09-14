package main

import "math/big"

// PayStrategy 策略接口
type PayStrategy interface {
	Pay(amount *big.Float)
}
