package main;

/**
 * 访问者模式演示 - 黑心小公司的项目分工
 * 
 * 访问者模式的核心思想：
 * 1. 将数据结构（员工）与操作（项目工作）分离
 * 2. 通过双分派机制实现不同类型员工在不同项目中的不同行为
 * 3. 符合开闭原则：新增项目类型无需修改员工类
 */
public class Main {

    public static void main(String[] args) {

        // 创建黑心小公司
        BlackHeartCompany company = new BlackHeartCompany("996创新科技有限公司");
        
        // 招聘员工（公司的数据结构）
        company.addEmployee(new Tester("小王"));
        company.addEmployee(new Developer("小李"));
        company.addEmployee(new HRBP("小张"));
        
        System.out.println();
        company.showCompanyInfo();
        
        // 创建不同的项目访问者（不同的算法）
        ProjectVisitor imProject = new IMProjectVisitor();
        ProjectVisitor ecommerceProject = new ECommerceProjectVisitor();
        
        // 演示访问者模式：同样的员工，不同的项目，不同的工作安排
        System.out.println("📱 三个月前，老板跟风做IM:");
        company.startProject(imProject, "仿微信IM");

        // 同时，访问者模式符合 OOP 的设计原则，在于它将数据结构与操作分离
        System.out.println("🛒 现在，老板又想做电商:");
        company.startProject(ecommerceProject, "仿淘宝电商");
        
        // 扩展
        // System.out.println("🎮 老板脑洞大开，又想做游戏了:");
        // ProjectVisitor gameProject = new GameProjectVisitor();
        // company.startProject(gameProject, "仿王者荣耀手游");
    }
}

/**
 * 游戏项目访问者 - 演示访问者模式的扩展性
 */
class GameProjectVisitor implements ProjectVisitor {
    
    @Override
    public void visit(Tester tester) {

        System.out.println("=== 游戏项目中的" + tester.getName() + " ===");
        System.out.println(tester.getName() + " 在游戏项目中的工作：游戏陪玩");
       
        tester.executeTest();
        
        System.out.println();
    }
    
    @Override
    public void visit(Developer developer) {
        
        System.out.println("=== 游戏项目中的" + developer.getName() + " ===");
        System.out.println(developer.getName() + " 在游戏项目中的工作：开发王者荣耀");
        
        developer.writeCode();
        
        System.out.println();
    }
    
    @Override
    public void visit(HRBP hrbp) {
        
        System.out.println("=== 游戏项目中的" + hrbp.getName() + " ===");
        System.out.println(hrbp.getName() + " 在游戏项目中的工作：游戏推广大使");
       
        hrbp.organizeActivity();

        System.out.println();
    }
}
