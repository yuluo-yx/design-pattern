package main;

// 政治敏感词过滤器
public class PoliticalFilter implements SensitiveFilter {
    
    @Override
    public String doFilter(String text, FilterChain chain) {
    
        String result = text.replaceAll("国家领导人", "[敏感词]");
        System.out.println("[PoliticalFilter] 处理后: " + result);
    
        return chain.doFilter(result);
    }
}
