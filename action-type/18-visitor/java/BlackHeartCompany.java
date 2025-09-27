package main;

import java.util.ArrayList;
import java.util.List;

/**
 * 黑心小公司类
 * 管理所有员工，并支持访问者模式遍历
 */
public class BlackHeartCompany {

    private List<Employee> employees;

    private String companyName;
    
    public BlackHeartCompany(String companyName) {
        this.companyName = companyName;
        this.employees = new ArrayList<>();
    }
    
    /**
     * 添加员工（招聘）
     */
    public void addEmployee(Employee employee) {

        employees.add(employee);

        System.out.println("恭喜 " + employee.getName() + "（" + employee.getPosition() + 
                         "）加入" + companyName + "大家庭！");
    }
    
    /**
     * 移除员工（裁员）
     */
    public void removeEmployee(Employee employee) {
        employees.remove(employee);
        System.out.println(employee.getName() + " 被优化了，祝工作顺利！");
    }
    
    /**
     * 启动项目 - 访问者模式的核心应用
     * 让所有员工接受项目访问者，分配不同的工作
     */
    public void startProject(ProjectVisitor projectVisitor, String projectName) {

        System.out.println("🚀 " + companyName + " 启动 " + projectName + " 项目！");
        System.out.println("老板：兄弟们，新项目来了，大家各司其职！");
        System.out.println("================================================");
        
        // 遍历所有员工，让他们接受项目访问者
        for (Employee employee : employees) {
            employee.accept(projectVisitor);
        }
        
        System.out.println("================================================");
        System.out.println("老板：很好，大家都有活干了！记住，deadline是一个月！");
        System.out.println();
    }
    
    /**
     * 获取员工总数
     */
    public int getEmployeeCount() {

        return employees.size();
    }
    
    /**
     * 显示公司概况
     */
    public void showCompanyInfo() {

        System.out.println("=== " + companyName + " 公司概况 ===");
        System.out.println("员工总数：" + employees.size() + " 人");
        System.out.println("员工列表：");

        for (Employee employee : employees) {
            System.out.println("- " + employee.getName() + "（" + employee.getPosition() + "）");
        }
        
        System.out.println();
    }
}