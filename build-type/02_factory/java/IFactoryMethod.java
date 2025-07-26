package main;

// 工厂方法

// 具体的产品实现
class FactoryMethodDashScope implements ChatModel {

    @Override
    public void call() {
        System.out.println("DashScope model called.");
    }
}

class FactoryMethodOpenAI implements ChatModel {

    @Override
    public void call() {
        System.out.println("OpenAI model called.");
    }
}

// 工厂接口
public interface IFactoryMethod {

    ChatModel createChatModel();
}

// 具体产品工厂，和产品一一对应
// 客户端通过 new DashScopeChatModelFactory().createChatModel() 或 new OpenAIChatModelFactory().createChatModel() 来获取具体的产品实例
// 进而调用 call 方法。

class DashScopeChatModelFactory implements IFactoryMethod {

    @Override
    public ChatModel createChatModel() {
        return new FactoryMethodDashScope();
    }
}

class OpenAIChatModelFactory implements IFactoryMethod {

    @Override
    public ChatModel createChatModel() {
        return new FactoryMethodOpenAI();
    }
}
