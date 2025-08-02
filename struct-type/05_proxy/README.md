# 代理（委托）模式 （Proxy Design Pattern）

代理模式：由于某些原因需要给原始对象提供一个代理以控制对该对象的访问。这时，访问对象不适合或者不能直接引用目标对象，代理对象作为访问对象和目标对象之间的中介。

## 1. 应用

1. OpenFeign的动态代理；
2. Spring AOP，事务等；
3. 符合开闭原则，可以在代理类中对原始类增加功能扩展，而不用修改原始类。负责为委托类预处理消息、过滤消息、把消息转发给委托类，以及事后对返回结果的处理等。代理类本身并不真正实现原始类服务，而是同过调用委托类的相关方法，来提供特定的服务。例如日志记录，事务处理，缓存等。

### 1.1 优点

1. 代理模式在客户端与目标对象之间起到一个中介作用和保护目标对象的作用；
2. 代理对象可以扩展目标对象的功能；
3. 代理模式能将客户端与目标对象分离，在一定程度上降低了系统的耦合度。

### 1.2 缺点

1. 增加系统复杂度；
2. 逻辑处理路径变长，处理速度变慢。

## 2. 实现分析

代理模式有多种实现，常见的是静态和动态代理两种模式。
其中动态代理在 Java 中常见用 JDK 和 CGLIB 代理。

以 Java 为例说明实现：

### 2.1 静态代理

```java
package main;

// 静态代理接口
interface IStaticProxy {

    // 检查是否需要调用原始类
    Boolean checkAccess();

    // 原始类调用前置处理
    // 默认方法，子类可以覆盖
    default void preHandle() {
        this.checkAccess();
    }

    // 原始类调用后置处理
    void postHandle();
}

// 静态代理类
public class StaticProxy implements IUserService, IStaticProxy {

    // 组合原始服务
    private final IUserService realService;

    public StaticProxy(IUserService realService) {
        this.realService = realService;
    }

    @Override
    public User getUser(String name) {

        this.preHandle();

        if (this.checkAccess()) {
            return this.realService.getUser(name);
        } else {
            System.out.println("Access denied.");
        }

        this.postHandle();

        return null;
    }

    @Override
    public Boolean checkAccess() {

        // 检查，这里省略，真实使用时需要根据业务场景判断
        return true;
    }

    @Override
    public void preHandle() {

        System.out.println("Pre-processing before service work.");
    }

    @Override
    public void postHandle() {

        System.out.println("Post-processing after service work.");
    }
}
```

### 2.2 动态代理

```java
package main;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

// JDK user 服务代理
// 此处只使用 JDK 动态代理实现
// 1. JDK 动态代理通过反射机制实现；代理对象没有实现接口时会抛出异常，CGLIB 动态代理通过 ASM 字节码实现。
// 2. JDK 动态代理中的被代理类必须实现至少一个接口；CGLIB 动态代理需要继承被代理类，所以不能代理 final 方法。
public class DynamicProxy {

    // 获取动态代理对象
    public static Object getDynamicProxy(Object targetObj) {

        ClassLoader classLoader = DynamicProxy.class.getClassLoader();
        Class<?>[] interfaces = targetObj.getClass().getInterfaces();

        InvocationHandler handler = new JDKUserProxyHandler(targetObj);

        // return dynamic_proxy_object
        return Proxy.newProxyInstance(classLoader, interfaces, handler);
    }
}

// User 服务接口代理逻辑处理器
class JDKUserProxyHandler implements InvocationHandler {

    // cache for users
    private Map<String, User> userCache;

    // target object to be proxied
    private Object targetObject;

    public JDKUserProxyHandler(Object targetObject) {
        this.targetObject = targetObject;
        this.userCache = new ConcurrentHashMap<>();
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {

        if (args.length > 0 && userCache.containsKey(args[0])) {

            System.out.println("Get from cache. key is: " + args[0]);
            return userCache.get(args[0]);
        }

        // 执行 targetObject.getUser()
        Object cacheUser = method.invoke(targetObject, args);

        // 添加到缓存
        // {"name":"user obj"}
        userCache.put((String) args[0],(User) cacheUser);

        return cacheUser;
    }
}
```

## 3. 面试题

Q：OpenFeign 中怎么样使用了代理设计模式？
<br>
A：OpenFeign的动态代理实现原理是使用 Java 动态代理技术，通过创建代理对象拦截 OpenFeign 接口的方法调用，并将这些调用转发到一个 HTTP 客户端进行处理。

Q：Spring AOP 代理实现
<br>
A：Spring 提供了 JDK 和 CGLIB 两种代理方式，默认使用 JDK 的动态代理来实现 AOP。**但是遵循一个原则：如果加入容器的目标对象有实现接口,用 JDK 代理；如果目标对象没有实现接口,用 CGLIB 代理。**
<br>
