package main

import (
	"fmt"
	"regexp"
)

// PoliticalFilter 政治敏感词过滤器

type PoliticalFilter struct{}

func (f *PoliticalFilter) Filter(text string, chain *FilterChain) string {

	result := text
	result = regexp.MustCompile(`国家领导人`).ReplaceAllString(result, "[敏感词]")
	fmt.Println("[PoliticalFilter] 处理后:", result)

	return chain.Filter(result)
}
