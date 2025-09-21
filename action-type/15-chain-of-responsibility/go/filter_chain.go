package main

import "fmt"

// SensitiveFilter 定义过滤器接口
// 每个过滤器实现 Filter 方法

type SensitiveFilter interface {
	Filter(text string, chain *FilterChain) string
}

// FilterChain 过滤器链，顺序执行所有过滤器

type FilterChain struct {
	filters []SensitiveFilter
	index   int
}

// AddFilter 添加过滤器到链
func (c *FilterChain) AddFilter(f SensitiveFilter) *FilterChain {

	c.filters = append(c.filters, f)
	return c
}

// Filter 执行链上的过滤器
func (c *FilterChain) Filter(text string) string {

	if c.index == len(c.filters) {
		fmt.Println("[FilterChain] 执行到链尾，过滤结束")
		return text
	}

	f := c.filters[c.index]
	c.index++
	fmt.Printf("[FilterChain] 执行第%d个过滤器: %T\n", c.index, f)

	return f.Filter(text, c)
}
