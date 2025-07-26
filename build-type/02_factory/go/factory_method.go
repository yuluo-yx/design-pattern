package main

import (
	"fmt"
)

// 工厂方法

type FactoryMethodDashScope struct {
}

func (f *FactoryMethodDashScope) Call() {

	fmt.Println("Factory Method, DashScope model Call")
}

type FactoryMethodOpenAi struct {
}

func (f *FactoryMethodOpenAi) Call() {

	fmt.Println("Factory Method, OpenAi model Call")
}

// IChatModel 工厂方法接口
type IChatModelFactoryMethod interface {

	// 建造方法，返回 IChatModel 的实现类
	createModel() IChatModel
}

type FactoryMethodDashScopeFactory struct {
}

// DashScope ChatModel 的工厂方法实现
func (f *FactoryMethodDashScopeFactory) createModel() IChatModel {
	return &FactoryMethodDashScope{}
}

type FactoryMethodOpenAiFactory struct {
}

// OpenAi ChatModel 的工厂方法实现
func (f *FactoryMethodOpenAiFactory) createModel() IChatModel {
	return &FactoryMethodOpenAi{}
}

// 使用简单工厂封装一下工厂方法
func NewChatModelFactoryMethod(types string) IChatModelFactoryMethod {
	switch types {
	case "dashscope":
		return &FactoryMethodDashScopeFactory{}
	case "openai":
		return &FactoryMethodOpenAiFactory{}
	default:
		return nil
	}
}
