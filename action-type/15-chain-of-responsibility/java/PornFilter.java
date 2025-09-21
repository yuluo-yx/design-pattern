package main;

// 涉黄涉暴过滤器
public class PornFilter implements SensitiveFilter {
   
    @Override
    public String doFilter(String text, FilterChain chain) {
        
        String result = text.replaceAll("黄色|暴力", "[敏感词]");
        System.out.println("[PornFilter] 处理后: " + result);
   
        return chain.doFilter(result);
    }
}
