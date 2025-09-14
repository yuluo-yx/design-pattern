package main;

import java.math.BigDecimal;

public class Main {
    
    public static void main(String[] args) {

        PayContext context = new PayContext();

        // 用户选择支付宝
        context.setStrategy(new AliPayStrategy());
        context.pay(new BigDecimal("100.00"));

        // 用户选择微信
        context.setStrategy(new WeChatPayStrategy());
        context.pay(new BigDecimal("200.00"));
    }
}
