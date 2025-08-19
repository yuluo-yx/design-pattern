# 5. 迪米特法则（最少知识原则）（Demeter Principle）

迪米特法则，又称为最少知识原则。是指一个对象应该对其他对象有尽可能少的了解。这意味着一个类应该只与直接的朋友（即与其直接相关的对象）进行交互，而不应该与“陌生人”进行交互，从而降低类之间的耦合度，提高系统的可维护性和可复用性。

## 5.1 优缺点

在 OOP 中有一些众所周知的抽象概念：封装、继承和多态。理论上可以用来设计出清晰的设计和良好的代码。
虽然这些都是非常重要的 OOP 概念，但是不够实用，不能直接用于系统设计和开发。这些概念是比较主观的，非常依赖于使用人的经验和知识。
对于其他概念，如单一职责原则、开闭原则等，也是一样的。

迪米特法则的独特之处在于它简洁而准确的定义，是 OOP 设计原则更精确的描述。允许在编写代码时直接应用，自动遵循了 OOP 的设计原则。**体现了高内聚和低耦合的理念。**

### 5.1.1 优点

1. 最直观的是降低了模块的耦合性，因为迪米特法则强调和直接相关的朋友对象交互，不与陌生人交互；
2. 不直接依赖于其他陌生人或陌生人交互，提升代码的可维护性。

### 5.1.2 缺点

1. 增加类的复杂性，因为不和其他陌生类交互，所以需要添加一些其他适配方法；
2. 过度使用迪米特法则可能导致过度设计，使得系统中存在过多的中间类和方法，反而降低了系统的可读性和可维护性。

## 5.2. 概念解释

上面中提到了 ”朋友“ 和 ”陌生人“。这两个词什么意思？

朋友：每个对象都会与其他对象之间都存在一定程度的耦合关系，其中主要耦合方式包含：依赖、关联、组合、聚合等等。如果某个类作为被调用者，在其调用方的类中的内部主要体现为：类中成员变量的类型、方法参数类型、方法返回值类型，那么该类就属于调用者的“朋友”。

陌生人：对于某个类而言，当它作为“被调用者”在“调用者”当中没有作为：类中成员变量的类型、方法参数类型、方法返回值类型，这些形式出现在类中，而是仅作为局部变量出现在某个方法体中，那么对于这种情况，该类就属于“陌生人”

## 5.3. 例子理解

抽象的概念往往难以理解，下面我们举个例子来理解下：现在有三类角色：A 公司的产品和程序员和用户 B 。

现在，用户 B 有一个小程序的需求，想找 A 公司来委托完成。于是找到公司责任人，派了产品经理来了解用户需求并制定方案。然后交由 A 公司的开发来完成。
但是用户 B 不放心全权委托给 A 公司，于是全程跟踪直到完成。过程中指指点点，这里改哪里改。因为用户 B 不是专业的，所以想法有些天马行空 + 不可实现。但是 A 
公司的程序员因为是客户所以不好拒绝。于是简单的小程序需求代码写的复杂且拖延了进度。导致 A 公司开发，产品和用户 B 都不爽。

这些问题直接导致了 A 公司没有发挥自己的特长，用户 B 没有看到想要的产品。用户 B 在整个过程中知道了太多软件开发的细节以及产品设计等等。这样只会增加软件开发的周期
和交付质量。增加了复杂程度。

实际上，用户 B 不需要关注太多的细节，只需要关注最后的产出即可。用迪米特法则分析。这里的 A 公司产品和用户 B 是朋友。用户 B 和 A 公司开发是陌生人。

于是 A 公司使用迪米特法则优化流程，用户 B 只和 A 公司产品对接，表达想要的软件产品需求，最后关注产出即可。降低终端用户和开发之间的依赖关系，避免各种各样的依赖问题。

## 5.4. 代码实现

### 5.4.1 不使用迪米特法则

```java
package main;

public class Test {
    public static void main(String[] args) {
        UserB userB = new UserB();
        userB.Todo();
    }
}

class UserB {
    private static final String todo = "小程序";

    public void Todo() {
        System.out.println("用户描述需求：" + todo);

        // 用户 B 联系到了 A 公司产品
        ACompanyPM pm = new ACompanyPM();
        
        // 产品理解需求
        pm.understand(todo);

        // A 公司产品根据用户描述产出需求文档
        String pmDocs = pm.make(todo);

        // A 公司程序员开始开发
        ACompanyDev developer = new ACompanyDev();
        
        // 用户B要求全程参与开发
        pm.dev(pmDocs, developer, this);
    }

    // 用户B干预开发过程
    public void interrupt() {
        System.out.println("用户B: 这里不对，要改!");
        System.out.println("用户B: 我想要加个新功能!");
    }
}

class ACompanyPM {
    public String make(String desc) {
        return "用户需求 " + desc + "产品文档";
    }

    public void understand(String desc) {
        System.out.println("理解用户需求: " + desc);
    }

    public void dev(String pmDocs, ACompanyDev developer, UserB user) {
        System.out.println("开始开发...");
        
        // 用户不断干预开发过程
        user.interrupt();
        
        // 被迫修改需求和代码
        System.out.println("修改需求和代码...");
        developer.make(pmDocs);
        
        System.out.println("开发进度延迟，代码复杂度增加...");
    }

    public void dev(String pmDocs, ACompanyDev developer) {
        System.out.println("正常开发流程...");
        developer.make(pmDocs);
    }
}

class ACompanyDev {
    public void make(String pmDocs) {
        System.out.println("按照产品文档" + pmDocs + "开发需求...");
    }
}
```

### 5.4.2 迪米特法则优化

```java
package main;

public class Test {
    public static void main(String[] args) {
        UserB userB = new UserB();
        userB.requestProduct();
    }
}

class UserB {
    private static final String todo = "小程序";

    public void requestProduct() {
        System.out.println("用户描述需求：" + todo);

        // 用户B只和产品经理交互
        ACompanyPM pm = new ACompanyPM();
        
        // 提出需求
        pm.collectRequirement(todo);
        
        // 等待产品完成
        String product = pm.deliverProduct();
        
        // 用户只关注最终产品
        System.out.println("收到产品: " + product);
    }
}

class ACompanyPM {
    private ACompanyDev developer;
    private String requirement;
    private String pmDocs;

    public ACompanyPM() {
        // 产品经理内部管理开发资源
        this.developer = new ACompanyDev();
    }

    public void collectRequirement(String desc) {
        System.out.println("理解并确认用户需求: " + desc);
        this.requirement = desc;
        
        // 转化为产品文档
        this.pmDocs = "需求" + desc + "的产品文档";
    }

    public String deliverProduct() {
        System.out.println("产品经理推进开发...");
        
        // 产品经理负责和开发对接
        developer.develop(pmDocs);
        
        // 产品经理负责质量把控
        verifyProduct();
        
        return "完成的" + requirement;
    }

    private void verifyProduct() {
        System.out.println("产品经理验证产品质量...");
    }
}

class ACompanyDev {
    public void develop(String pmDocs) {
        System.out.println("根据产品文档开发: " + pmDocs);
    }
}
```

这里只是一个简单的例子实现，在实际的开发中，可能会运用抽象类，接口等实现。在设计模式中，门面模式和中介者模式体现了迪米特法则的思想。
