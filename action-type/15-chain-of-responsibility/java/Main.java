package main;

public class Main {

	public static void main(String[] args) {
	
        String text = "我的手机号是13812345678，国家领导人，黄色暴力";
		System.out.println("原始文本: " + text);

		FilterChain chain = new FilterChain();
		chain.addFilter(new PersonalInfoFilter())
			 .addFilter(new PoliticalFilter())
			 .addFilter(new PornFilter());

		String result = chain.doFilter(text);
	
        System.out.println("最终过滤结果: " + result);
	}
}
