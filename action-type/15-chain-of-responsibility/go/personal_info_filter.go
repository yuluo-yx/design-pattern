package main

import (
	"fmt"
	"regexp"
)

// PersonalInfoFilter 个人信息过滤器

type PersonalInfoFilter struct{}

func (f *PersonalInfoFilter) Filter(text string, chain *FilterChain) string {

	// 手机号简单替换
	re := regexp.MustCompile(`\d{11}`)
	result := re.ReplaceAllString(text, "[手机号]")
	fmt.Println("[PersonalInfoFilter] 处理后:", result)

	return chain.Filter(result)
}
