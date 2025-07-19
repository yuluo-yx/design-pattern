package main;

public class SingletonLazy {

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

}
