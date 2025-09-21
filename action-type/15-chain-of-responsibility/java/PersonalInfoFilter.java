package main;

// 个人信息过滤器
public class PersonalInfoFilter implements SensitiveFilter {
    
    @Override
    public String doFilter(String text, FilterChain chain) {
    
        String result = text.replaceAll("[0-9]{11}", "[手机号]");
        System.out.println("[PersonalInfoFilter] 处理后: " + result);
    
        return chain.doFilter(result);
    }
}
