package main

func main() {

	// 简单工厂
	dashscopeModel := NewChatModel("dashscope")
	dashscopeModel.Call()
	openAiModel := NewChatModel("openai")
	openAiModel.Call()

	// 工厂方法
	dashscopeFactory := NewChatModelFactoryMethod("dashscope")
	dashscopeFactory.createModel().Call()
	openAiFactory := NewChatModelFactoryMethod("openai")
	openAiFactory.createModel().Call()

	// 抽象工厂
	dashScopeAbstractFactory := NewChatModelAbstractFactory("dashscope")
	dashScopeAbstractFactory.CreateChatOptions().Options()
	dashScopeAbstractFactory.CreateApiKey().Key()
	openAiAbstractFactory := NewChatModelAbstractFactory("openai")
	openAiAbstractFactory.CreateChatOptions().Options()
	openAiAbstractFactory.CreateApiKey().Key()

}
