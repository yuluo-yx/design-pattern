package main;

public class SingletonEager {

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

}
