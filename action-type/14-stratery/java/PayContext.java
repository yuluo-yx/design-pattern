package main;

import java.math.BigDecimal;

public class PayContext {

    // ctx 持有的支付策略
    private PayStrategy strategy;

    public void setStrategy(PayStrategy strategy) {
        this.strategy = strategy;
    }

    // ctx 中的执行支付方法
    public void pay(BigDecimal amount) {
        
        if (strategy == null) {
            throw new IllegalStateException("未设置支付策略");
        }
        
        strategy.pay(amount);
    }
}
