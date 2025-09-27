package main;

/**
 * 开发员工类 - 访问者的具体元素
 */
public class Developer implements Employee {

    private String name;
    
    public Developer(String name) {
        this.name = name;
    }
    
    @Override
    public void accept(ProjectVisitor visitor) {

        // 双分派：这里的 this 是 Developer 类型
        visitor.visit(this);
    }
    
    @Override
    public String getName() {

        return name;
    }
    
    @Override
    public String getPosition() {

        return "开发工程师";
    }
    
    /**
     * 开发员工特有的方法 - 编写代码
     */
    public void writeCode() {

        System.out.println(name + " 正在疯狂敲代码...");
    }
    
    /**
     * 开发员工特有的方法 - 修复Bug
     */
    public void fixBug() {

        System.out.println(name + " 修复了一个Bug（可能引入了两个新Bug）");
    }
}
