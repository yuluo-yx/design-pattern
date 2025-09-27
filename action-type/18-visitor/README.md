# 访问者模式（Visitor Design Pattern）

访问者模式是一种行为设计模式， 它能将算法与其所作用的对象隔离开来。

> 实际应用并不多。

## 使用场景

1. 如果需要对一个复杂对象结构（例如对象树）中的所有元素执行某些操作；
2. 当某个行为仅在类层次结构中的一些类中有意义，而在其他类中没有意义时。

## 伪动态双分派

访问者模式中有一种称之为伪动态双分派的技术应用。这只在 OOP 类型中存在，Go 中并没有类似的技术手段。

因此见得设计模式并不是适用于所有语言，不同阶段。而是需要根据不同场景选择。

### 分派

首先看一个 Java Map 的例子：

Map map = new HashMap<>();

其中 Map 被称为变量 map 的静态类型，因为在编译时期就确定了。而变量 map 的实际类型是 Map 的子类 HashMap，这在运行时确定。另一种叫法是 Map 是明显类型，HashMap 是实际类型。那么所谓分派就是：

- 静态分派 Static Dispatch：发生在编译时期，分派根据静态类型信息发生。方法重载就是静态分派。
- 动态分派 Dynamic Dispatch：发生在运行时期，动态分派动态地置换掉某个方法。Java 通过方法的重写支持动态分派。

### 伪动态双分派

双分派就是在选择一个方法的时候，不仅仅要根据消息接收者（receiver）的运行时区别，还要根据参数的运行时区别。**双分派实现动态绑定的本质，就是在重载方法委派的前面加上了继承体系中覆盖的环节，由于覆盖是动态的，所以重载就是动态的。**

1. **第一次分派**：通过方法重写（覆盖）实现，在运行时根据对象的实际类型选择方法；
2. **第二次分派**：通过方法重载实现，根据参数类型选择具体的重载版本。

#### Java代码示例

```java
// 访问者接口
interface Visitor {
    void visit(ConcreteElementA element);
    void visit(ConcreteElementB element);
}

// 元素接口
interface Element {
    void accept(Visitor visitor);
}

// 具体元素A
class ConcreteElementA implements Element {
    @Override
    public void accept(Visitor visitor) {
        // 第一次分派：运行时根据visitor的实际类型调用相应的visit方法（多态）
        // 第二次分派：编译时根据this的静态类型(ConcreteElementA)选择重载的visit方法
        visitor.visit(this); // this的静态类型是ConcreteElementA
    }
}

// 具体元素B
class ConcreteElementB implements Element {
    @Override
    public void accept(Visitor visitor) {
        visitor.visit(this); // this的静态类型是ConcreteElementB
    }
}

// 具体访问者1
class ConcreteVisitor1 implements Visitor {
    @Override
    public void visit(ConcreteElementA element) {
        System.out.println("ConcreteVisitor1 访问 ConcreteElementA");
    }
    
    @Override
    public void visit(ConcreteElementB element) {
        System.out.println("ConcreteVisitor1 访问 ConcreteElementB");
    }
}

// 具体访问者2
class ConcreteVisitor2 implements Visitor {
    @Override
    public void visit(ConcreteElementA element) {
        System.out.println("ConcreteVisitor2 访问 ConcreteElementA");
    }
    
    @Override
    public void visit(ConcreteElementB element) {
        System.out.println("ConcreteVisitor2 访问 ConcreteElementB");
    }
}

// 演示双分派
public class DoubleDispatchDemo {

    public static void main(String[] args) {
        // 创建元素数组（静态类型都是Element）
        Element[] elements = {
            new ConcreteElementA(),  // 实际类型是ConcreteElementA
            new ConcreteElementB()   // 实际类型是ConcreteElementB
        };
        
        // 创建访问者数组（静态类型都是Visitor）
        Visitor[] visitors = {
            new ConcreteVisitor1(),  // 实际类型是ConcreteVisitor1
            new ConcreteVisitor2()   // 实际类型是ConcreteVisitor2
        };
        
        // 双分派演示
        for (Element element : elements) {
            for (Visitor visitor : visitors) {
                // 第一次分派：根据element的运行时类型调用相应的accept方法
                // 第二次分派：在accept方法内部，根据visitor的运行时类型和this的编译时类型选择visit方法
                element.accept(visitor);
            }
        }
        
        /* 输出：
         * ConcreteVisitor1 访问 ConcreteElementA
         * ConcreteVisitor2 访问 ConcreteElementA  
         * ConcreteVisitor1 访问 ConcreteElementB
         * ConcreteVisitor2 访问 ConcreteElementB
         */
    }
}

#### 双分派的执行过程解析

1. **调用 `element.accept(visitor)`**：
   - 第一次分派：根据 `element` 的运行时类型（ConcreteElementA 或 ConcreteElementB）选择相应的 `accept` 方法

2. **在 `accept` 方法中调用 `visitor.visit(this)`**：
   - 第二次分派：根据 `visitor` 的运行时类型选择具体的访问者实现，同时根据 `this` 的编译时类型选择重载的 `visit` 方法

这样就实现了根据两个对象的类型来选择方法的双分派机制。

## 1. 应用

### 1.1 优点

1. 符合设计原则
2. 访问者对象可以在与各种对象交互时收集一些有用的信息。当想要遍历一些复杂的对象结构（例如对象树），并在结构中的每个对象上应用访问者时，这些信息可能会有所帮助。

### 1.2 缺点

1. 每次在元素层次结构中添加或移除一个类时，你都要更新所有的访问者。
2. 在访问者同某个元素进行交互时，它们可能没有访问元素私有成员变量和方法的必要权限。

## 2. 代码分析

举个例子：黑心小公司，只招了三个人：测试（Tester），开发（Developer）和一个人事（HRBP）。公司老板不仅想做养老系统还想做电商，还想做 IM。现在以 IM 和电商系统为例。

3个月前，老板跟风做 IM，给开发一个月时间开发一套微信 App，开发出来了，人力来发文宣传推广，测试也不能闲着，当陪聊？

现在，老板想做电商系统，给开发一个月时间做一个淘宝 App，人力电商直播带货，测试当客服。

这个例子中对象和算法分别是：

Tester: 陪聊，客服

Developer：淘宝开发，IM 开发

HRBP：运营，主播

现在开始写代码，参考同目录下的 java & go 文件夹。
