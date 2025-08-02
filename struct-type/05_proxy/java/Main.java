package main;

public class Main {

    public static void main(String[] args) {

        // 静态代理示例
        System.out.println("========================== 静态代理 ==========================");
        // 原始实现类（需要被代理的类）
        IUserService realService = new ServiceImpl();
        // 静态代理类（中间/委托人）委托访问原始类
        StaticProxy staticProxy = new StaticProxy(realService);
        staticProxy.getUser("JDK");
        System.out.println("========================== 静态代理 ==========================");

        // 动态代理示例
        System.out.println("========================== 动态代理 ==========================");
        // 原始实现类
        IUserService dpRealService = new ServiceImpl();
        // 动态代理获取中介对象
        IUserService dynamicProxyObj = (IUserService) DynamicProxy.getDynamicProxy(dpRealService);
        // 通过代理类调用原始对象的方法
        dynamicProxyObj.getUser("JDK");
        dynamicProxyObj.getUser("JDK");
        System.out.println("========================== 动态代理 ==========================");
    }

}
