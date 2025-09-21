package main;

public class Main {
	
    public static void main(String[] args) {
		
        LeaveRequestContext ctx = new LeaveRequestContext("张三", "家中有事请假2天");
		System.out.println("--- 发起请假 ---");
		ctx.submit();
		System.out.println("--- 主管审批同意 ---");
		ctx.approve();
		System.out.println("--- HRBP 审批同意 ---");
		ctx.approve();
		System.out.println("--- 尝试再次提交 ---");
		ctx.submit();

		System.out.println("--- 新流程：发起请假，主管拒绝 ---");
		LeaveRequestContext ctx2 = new LeaveRequestContext("李四", "身体不适请假1天");
		ctx2.submit();
		ctx2.reject();

		System.out.println("--- 新流程：发起请假，撤销 ---");
		LeaveRequestContext ctx3 = new LeaveRequestContext("王五", "家中有事请假3天");
		ctx3.cancel();
	}
}
