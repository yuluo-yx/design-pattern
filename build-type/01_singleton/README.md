# 单例设计模式（Singleton Design Pattern）

在系统启动到停止，只存在一个这样的对象（实例），这个类就是一个单例类。

## 1. 应用

在 Spring 中默认所有创建的 Bean 都是单例模式。

### 1.1 优点

1. 避免常见过多的对象实例，造成资源浪费，例如数据库连接对象可能是单例模式；

### 1.2 缺点

1. 隐藏类之间的依赖关系，对外提供唯一可访问的方法；
2. 对代码扩展性不友好；
3. 不支持有参数的构造函数；
4. 不方便 mock 进行测试。

## 2. 实现分析

实现分析以 Java 为例。

### 2.1 饿汉式

饿汉式：一个害怕饿的人总是很一开始就带很多食物。顾名思义，在系统启动就已经初始化了对象实例。

```java
// 静态变量加载顺序
// 1. 静态变量
// 2. 静态代码块
// 3. 静态方法
private static final Models instance = new Models();

private SingletonEager() {
}

public static Models getInstance() {

    return instance;
}
```

### 2.2 懒汉式

懒汉式：一个很懒的人会在饿了的时候才去寻找食物。在系统调用用到了才去初始化对象实例。

```java
// 在普通 new 对象过程中，会发生以下三个步骤：
// 1. 分配内存空间
// 2. 初始化对象
// 3. 将引用指向内存空间
// 但实际可能会发生指令重排序：
// 1. 分配内存空间
// 3. 将引用指向内存空间
// 2. 初始化对象 （此时可能返回未完全初始化的对象）
// 此时，使用 volatile 保证可见性，避免 jvm 动作
private static volatile Models obj;

private SingletonLazy() {
}

// 双重检查机制，线程安全
public static Models getLazyInstance1() {
    // 1. 检查对象是否已经创建, 避免多余的同步开销
    if (obj == null) {
        synchronized (Models.class) {
            // 2. 再次检查对象是否已经创, 避免重复创建
            if (obj == null) {
                obj = new Models();
            }
        }
    }
    return obj;
}

// Java 的另一种实现方式，静态内部类
// 只有在首次访问 SingletonHolder.INSTANCE 时，才会加载静态内部类
private static class SingletonHolder {

    private static final Models INSTANCE = new Models();
}

public static Models getLazyInstance2() {

    return SingletonHolder.INSTANCE;
}
```

> Tips: Java 中的 enum 默认单例，因此也常被用做单例模式设计。

## 3. 相关面试题

Q: Spring 中的 Controller 是单例还是多例，并发安全吗？
<br>
A: Controller 默认是单例的.
<br>
  1. 并且在方法中不要使用非静态的成员变量;
  2. 必须要定义一个非静态成员变量时，通过注解 `@Scope(“prototype”)`，设置为多例模式;
  3. 在 Controller 中使用 ThreadLocal 变量。

## 4. 扩展

1. Java 的初始化时机：https://blog.csdn.net/sinat_36817189/article/details/93928062
2. https://github.com/mohuishou/go-design-pattern/blob/master/01_singleton/readme.md
