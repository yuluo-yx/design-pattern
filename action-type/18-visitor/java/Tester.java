package main;

/**
 * 测试员工类 - 具体元素
 */
public class Tester implements Employee {
    private String name;
    
    public Tester(String name) {
        this.name = name;
    }
    
    @Override
    public void accept(ProjectVisitor visitor) {
        // 双分派的关键：这里的this是Tester类型
        visitor.visit(this);
    }
    
    @Override
    public String getName() {
        return name;
    }
    
    @Override
    public String getPosition() {
        return "测试工程师";
    }
    
    /**
     * 测试员工特有的方法 - 执行测试
     */
    public void executeTest() {
        System.out.println(name + " 正在执行测试用例...");
    }
    
    /**
     * 测试员工特有的方法 - 提交Bug报告
     */
    public void submitBugReport() {
        System.out.println(name + " 提交了Bug报告");
    }
}