package main;

/**
 * 电商项目访问者 - 具体访问者
 * 实现电商项目中各种员工的工作分配
 */
public class ECommerceProjectVisitor implements ProjectVisitor {
    
    @Override
    public void visit(Tester tester) {

        System.out.println("=== 电商项目中的" + tester.getName() + " ===");
        System.out.println(tester.getName() + " 在电商项目中的工作：客服小姐姐");

        tester.submitBugReport();

        System.out.println();
    }
    
    @Override
    public void visit(Developer developer) {

        System.out.println("=== 电商项目中的" + developer.getName() + " ===");
        System.out.println(developer.getName() + " 在电商项目中的工作：开发淘宝App");

        developer.writeCode();
        developer.fixBug();
        
        System.out.println();
    }
    
    @Override
    public void visit(HRBP hrbp) {
        
        System.out.println("=== 电商项目中的" + hrbp.getName() + " ===");
        System.out.println(hrbp.getName() + " 在电商项目中的工作：直播带货主播");
       
        hrbp.recruit(); 
        
        System.out.println();
    }
}