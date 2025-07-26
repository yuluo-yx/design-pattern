package main;

// 抽象工厂

// 举例：一个 ChatModel 包含 ChatOptions 和 ApiKey

interface ChatOptions {

    void options();
}

// DashScope 的 ChatOptions 实现
class AbstractChatOptionsDashScope implements ChatOptions {

    @Override
    public void options() {
        System.out.println("DashScope chat options configured.");
    }
}

// OpenAI 的 ChatOptions 实现
class AbstractChatOptionsOpenAI implements ChatOptions {

    @Override
    public void options() {
        System.out.println("OpenAI chat options configured.");
    }
}

interface ApiKey {

    void key();
}

// DashScope 的 ApiKey 实现
class AbstractApiKeyDashScope implements ApiKey {

    @Override
    public void key() {
        System.out.println("DashScope api key configured.");
    }
}

// OpenAI 的 ApiKey 实现
class AbstractApiKeyOpenAI implements ApiKey {

    @Override
    public void key() {
        System.out.println("OpenAI api key configured.");
    }
}

// 抽象工厂接口，将同一个产品的不同组成子产品封装成抽象工厂接口
// 而不需要单独每一个提供工厂子类
public interface IAbstractFactory {

    ChatOptions createChatOptions();
    ApiKey createApiKey();
}

class AbstractFactoryDashScope implements IAbstractFactory {

    @Override
    public ChatOptions createChatOptions() {
        return new AbstractChatOptionsDashScope();
    }

    @Override
    public ApiKey createApiKey() {
        return new AbstractApiKeyDashScope();
    }
}

class AbstractFactoryOpenAI implements IAbstractFactory {

    @Override
    public ChatOptions createChatOptions() {
        return new AbstractChatOptionsOpenAI();
    }

    @Override
    public ApiKey createApiKey() {
        return new AbstractApiKeyOpenAI();
    }
}
