package main;

// 简单工厂

// 产品的不同实现提供方
class SimpleDashScope implements ChatModel {
    @Override
    public void call() {
        System.out.println("DashScope model called.");
    }
}

class SimpleOpenAI implements ChatModel {
    @Override
    public void call() {
        System.out.println("OpenAI model called.");
    }
}

// 简单工厂类实现
public class SimpleFactory {

    private SimpleFactory() {
        // 私有构造函数，防止直接实例化简单工厂
    }

    // 根据传入参数的不同构建不同的产品实例
    public static ChatModel create(String modelType) {

        // use switch-case or if-else to determine the model type
        if ("DashScope".equalsIgnoreCase(modelType)) {
            return new SimpleDashScope();
        } else if ("OpenAI".equalsIgnoreCase(modelType)) {
            return new SimpleOpenAI();
        } else {
            throw new IllegalArgumentException("Unknown model type: " + modelType);
        }
    }

}
