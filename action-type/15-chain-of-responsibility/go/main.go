package main

import "fmt"

func main() {

	text := "我的手机号是13812345678，国家领导人，黄色暴力"
	fmt.Println("原始文本:", text)

	chain := &FilterChain{}
	chain.AddFilter(&PersonalInfoFilter{}).
		AddFilter(&PoliticalFilter{}).
		AddFilter(&PornFilter{})

	result := chain.Filter(text)

	fmt.Println("最终过滤结果:", result)
}
