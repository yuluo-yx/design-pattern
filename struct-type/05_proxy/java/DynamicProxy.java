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
