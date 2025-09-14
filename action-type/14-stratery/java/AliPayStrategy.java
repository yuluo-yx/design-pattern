package main;

import java.math.BigDecimal;

public class AliPayStrategy implements PayStrategy {

    @Override
    public void pay(BigDecimal amount) {
        
        System.out.println("使用支付宝支付：" + amount + "元");
    }
}
