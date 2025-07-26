package main;

public class Main {

    public static void main(String[] args) {

        // 简单工厂使用
        ChatModel simpleDashScope = SimpleFactory.create("DashScope");
        simpleDashScope.call();
        ChatModel simpleOpenAI = SimpleFactory.create("OpenAI");
        simpleOpenAI.call();

        // 工厂方法
        ChatModel factoryMethodDashScope = new DashScopeChatModelFactory().createChatModel();
        factoryMethodDashScope.call();
        ChatModel factoryMethodOpenAI = new OpenAIChatModelFactory().createChatModel();
        factoryMethodOpenAI.call();

        // 抽象工厂
        IAbstractFactory abstractFactoryDashScope = new AbstractFactoryDashScope();
        ChatOptions chatOptionsDashScope = abstractFactoryDashScope.createChatOptions();
        ApiKey apiKeyDashScope = abstractFactoryDashScope.createApiKey();
        chatOptionsDashScope.options();
        apiKeyDashScope.key();

        IAbstractFactory abstractFactoryOpenAI = new AbstractFactoryOpenAI();
        ChatOptions chatOptionsOpenAI = abstractFactoryOpenAI.createChatOptions();
        ApiKey apiKeyOpenAI = abstractFactoryOpenAI.createApiKey();
        chatOptionsOpenAI.options();
        apiKeyOpenAI.key();

    }

}
