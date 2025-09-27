package main;

/**
 * 员工接口 - 元素接口
 * 定义了接受项目访问者的方法
 */
public interface Employee {
    
    /**
     * 接受项目访问者
     * @param visitor 项目访问者
     */
    void accept(ProjectVisitor visitor);
    
    /**
     * 获取员工姓名
     */
    String getName();
    
    /**
     * 获取员工职位
     */
    String getPosition();
}