package main;

// 过滤器接口，定义处理方法
public interface SensitiveFilter {
    
    /**
     * @param text 输入文本
     * @return 过滤后的文本
     */
    String doFilter(String text, FilterChain chain);

}
