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
