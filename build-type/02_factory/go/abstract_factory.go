package main

import (
	"fmt"
)

// 抽象工厂实现

// 举例：一个 ChatModel 包含 ChatOptions 和 ApiKey

type ChatOptions interface {
	Options()
}

type DashScopeChatOptions struct {
}

func (d *DashScopeChatOptions) Options() {
	fmt.Println("DashScope ChatOptions")
}

type OpenAiChatOptions struct {
}

func (o *OpenAiChatOptions) Options() {
	fmt.Println("OpenAi ChatOptions")
}

type ApiKey interface {
	Key()
}

type DashScopeApiKey struct {
}

func (d *DashScopeApiKey) Key() {
	fmt.Println("DashScope ApiKey")
}

type OpenAiApiKey struct {
}

func (o *OpenAiApiKey) Key() {

	fmt.Println("OpenAi ApiKey")
}

// 抽象工厂接口，将同一个产品的不同组成子产品封装成抽象工厂接口
// 而不需要单独每一个提供工厂子类
type IChatModelAbstractFactory interface {
	// 创建 ChatOptions
	CreateChatOptions() ChatOptions
	// 创建 ApiKey
	CreateApiKey() ApiKey
}

type DashScopeChatModelFactory struct {
}

func (d *DashScopeChatModelFactory) CreateChatOptions() ChatOptions {
	return &DashScopeChatOptions{}
}

func (d *DashScopeChatModelFactory) CreateApiKey() ApiKey {

	return &DashScopeApiKey{}
}

type OpenAiChatModelFactory struct {
}

func (o *OpenAiChatModelFactory) CreateChatOptions() ChatOptions {
	return &OpenAiChatOptions{}
}

func (o *OpenAiChatModelFactory) CreateApiKey() ApiKey {

	return &OpenAiApiKey{}
}

// 抽象工厂基于工厂方法，用一个简单工厂封装工厂方法
// 只不过一个工厂方法可创建多个相关的类
func NewChatModelAbstractFactory(types string) IChatModelAbstractFactory {
	switch types {
	case "dashscope":
		return &DashScopeChatModelFactory{}
	case "openai":
		return &OpenAiChatModelFactory{}
	default:
		return nil
	}
}
