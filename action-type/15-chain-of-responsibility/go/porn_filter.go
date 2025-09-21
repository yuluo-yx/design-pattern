package main

import (
	"fmt"
	"regexp"
)

// PornFilter 涉黄涉暴过滤器

type PornFilter struct{}

func (f *PornFilter) Filter(text string, chain *FilterChain) string {

	result := regexp.MustCompile(`黄色|暴力`).ReplaceAllString(text, "[敏感词]")
	fmt.Println("[PornFilter] 处理后:", result)

	return chain.Filter(result)
}
