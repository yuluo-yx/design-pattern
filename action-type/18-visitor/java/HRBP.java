package main;

/**
 * HR员工类 - 具体元素
 */
public class HRBP implements Employee {

    private String name;
    
    public HRBP(String name) {
        this.name = name;
    }
    
    @Override
    public void accept(ProjectVisitor visitor) {
        // 双分派：这里的 this 是 HRBP 类型
        visitor.visit(this);
    }
    
    @Override
    public String getName() {

        return name;
    }
    
    @Override
    public String getPosition() {

        return "人力资源业务伙伴";
    }
    
    /**
     * HR员工特有的方法 - 招聘
     */
    public void recruit() {

        System.out.println(name + " 正在招聘新员工（画饼中...）");
    }
    
    /**
     * HR员工特有的方法 - 组织活动
     */
    public void organizeActivity() {
        
        System.out.println(name + " 组织了团建活动");
    }
}