package main;

import java.util.ArrayList;
import java.util.List;

// 过滤器链，按顺序执行所有过滤器
public class FilterChain {

    private final List<SensitiveFilter> filters = new ArrayList<>();
    
    private int index = 0;

    public FilterChain addFilter(SensitiveFilter filter) {
        
        filters.add(filter);
        return this;
    }

    public String doFilter(String text) {
        
        if (index == filters.size()) {
            System.out.println("[FilterChain] 执行到链尾，过滤结束");
            return text;
        }
        
        SensitiveFilter filter = filters.get(index++);
        
        System.out.println("[FilterChain] 执行第" + index + "个过滤器: " + filter.getClass().getSimpleName());
        
        return filter.doFilter(text, this);
    }
}
