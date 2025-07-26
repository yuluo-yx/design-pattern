# 工厂模式 （Factory Design Pattern）

在面向对象编程种，任何事物都是一个对象。一般在 Java 中会通过 new 关键字来创建对象实例。
但是通过 new 关键字难免需要考虑不同子类对象，在对象构造之前做参数准备等，其最大的目的是`解耦合`。
因此工厂设计模式提倡使用工厂类来创建对象。

例如：

```java
// 同一个接口的不同实现
ChatModel a = new DashScopeChatModel();
ChatModel b = new OpenAiChatModel();

// 解耦合
在 controller 中注入 service 接口的实现类，将 service 和 controller 解耦。
```

Spring 中 `BeanPostProcessor` 接口，提供 Spring Bean 初始化前后的操作时机等。

## 1. 应用

在 Spring 框架中，大量使用了工厂设计模式，例如 `BeanFactory` 和 `ApplicationContext` 以及它们的实现 `XmlBeanFactory` 和 `DefaultListableBeanFactory` 等。

## 2. 工厂模式实现

在 《设计模式》 书中，将工厂方法分为两类，一类是工厂方法模式，一类是抽象工厂模式。
将简单工厂看作是工厂方法的一种特例（静态工厂方法模式），归于工厂方法中。

### 2.1 简单工厂模式

> 以 Java 实现为例。

```java
// 简单工厂

// 抽象产品接口
interface ChatModel {
    void call();
}

// 产品的不同实现提供方
class DashScope implements ChatModel {
    @Override
    public void call() {
        System.out.println("DashScope model called.");
    }
}

class OpenAI implements ChatModel {
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

        if ("DashScope".equalsIgnoreCase(modelType)) {
            return new DashScope();
        } else if ("OpenAI".equalsIgnoreCase(modelType)) {
            return new OpenAI();
        } else {
            throw new IllegalArgumentException("Unknown model type: " + modelType);
        }
    }

}
```

缺点：新增一个产品实现时，需要更新工厂方法。不符合开闭原则。

### 2.2 工厂方法模式

工厂方法(Factory Method)模式，又称多态工厂(Polymorphic Factory)模式或虚拟构造器(Virtual Constructor)模式。工厂方法模式通过定义工厂抽象父类(或接口)负责定义创建对象的公共接口，而工厂子类(或实现类)则负责生成具体的对象。

```java
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
```

工厂方法模式的优点如下：

- 客户端只依赖产品的抽象，符合依赖倒置原则；无需关注具体产品，符合迪米特(最少知识)原则。
- 工厂方法模式符合开放封闭原则，新增产品时只需增加相应的产品和工厂类，而无需修改现有代码。
- 工厂方法模式符合单一职责原则，每个具体的工厂类只负责创建对应的产品。
- 工厂方法模式不使用静态工厂方法，可以形成基于继承的等级结构。

缺点如下：

- 新增产品时除增加新产品类外，还要增加对应的具体工厂类，没有简单工厂代码简洁、高效。
- 为了扩展性而进一步引入抽象层，增加了系统的抽象性和理解难度。

### 2.3 抽象工厂模式

抽象工厂模式由工厂接口(或抽象类)、一组实现工厂接口工厂类、一个产品接口(或抽象类)和一组实现产品接口的具体产品组成。
和工厂方法大致相同，不同的是，抽象工厂会将同一产品的多个组成部分抽象到一个工厂方法中进行创建，而不需要多个工厂子类。

在实际开发中应用并不多，Spring 中主要以简单工厂和工厂方法居多。

## 3. 相关面试题

Q：什么是工厂方法/说说工厂方法定义

A: 工厂模式是一种创建型设计模式
   定义一个用于创建对象的接口，把对象的创建和使用的过程分开，解耦合。

Q：能手写一下简单工厂吗？

A：见代码

Q：工厂方法在 Spring 的用法

A：BeanFactory

## 4. 扩展

IOC（Inversion of Control） 与 DI：

DI （Dependency Injection） 容器是 Spring 中的核心特性，用来管理对象创建，生命周期，依赖关系等。可以看作是工厂方法的高级用法。
IOC：控制反转，指的是将设计好的实例 Bean 交由 DI 容器创建，依赖对象的获取被反转了。（正转既是由开发者手动控制）

两者关系：两者总是同时出现，IOC 是一种设计思想，目的在于解除对象之间的依赖关系，松耦合。而 DI 是实现 IOC 的具体实现方式，通过 IOC 容器，将对象依赖的其他对象注入到该对象中。
