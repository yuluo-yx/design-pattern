package main

import (
	"fmt"
)

// 简单工厂

type SimpleDashScope struct {
}

func (s SimpleDashScope) Call() {

	fmt.Println("Simple Factory, DashScope model Call")
}

type SimpleOpenAi struct {
}

func (s SimpleOpenAi) Call() {
	fmt.Println("Simple Factory, OpenAi model Call")
}

func NewChatModel(types string) IChatModel {
	switch types {
	case "dashscope":
		return SimpleDashScope{}
	case "openai":
		return SimpleOpenAi{}
	default:
		return nil
	}
}
