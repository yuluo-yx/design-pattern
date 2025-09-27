package main;

/**
 * IM项目访问者 - 具体访问者
 * 实现IM项目中各种员工的工作分配
 */
public class IMProjectVisitor implements ProjectVisitor {
    
    @Override
    public void visit(Tester tester) {

        System.out.println("=== IM项目中的" + tester.getName() + " ===");
        System.out.println(tester.getName() + " 在IM项目中的工作：当陪聊机器人");
        
        tester.executeTest();

        System.out.println();
    }
    
    @Override
    public void visit(Developer developer) {
        
        System.out.println("=== IM项目中的" + developer.getName() + " ===");
        System.out.println(developer.getName() + " 在IM项目中的工作：开发微信App");
 
        developer.writeCode();
       
        System.out.println();
    }
    
    @Override
    public void visit(HRBP hrbp) {
        
        System.out.println("=== IM项目中的" + hrbp.getName() + " ===");
        System.out.println(hrbp.getName() + " 在IM项目中的工作：宣传运营");

        hrbp.organizeActivity();
        
        System.out.println();
    }
}