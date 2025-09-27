package main;

/**
 * 项目访问者接口
 * 定义了对不同员工类型的访问操作
 */
public interface ProjectVisitor {
    /**
     * 访问测试员工
     */
    void visit(Tester tester);
    
    /**
     * 访问开发员工
     */
    void visit(Developer developer);
    
    /**
     * 访问HR员工
     */
    void visit(HRBP hrbp);
}