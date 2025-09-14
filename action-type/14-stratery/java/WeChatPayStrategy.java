package main;

import java.math.BigDecimal;

public class WeChatPayStrategy implements PayStrategy {

    @Override
    public void pay(BigDecimal amount) {
        
        System.out.println("使用微信支付：" + amount + "元");
    }
}
