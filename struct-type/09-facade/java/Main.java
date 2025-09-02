package main;

public class Main {

    public static void main(String[] args) {
        // 创建门面
        Facade shoppingFacade = new Facade();
        
        // 客户端只需要调用一个简单的方法即可完成整个购物流程
        System.out.println("=== 开始购物流程 ===");
        shoppingFacade.placeOrder();
        System.out.println("=== 购物流程完成 ===");
    }
    
}
