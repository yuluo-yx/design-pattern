package main;

// 主管审批状态
public class ManagerApprovalState implements LeaveRequestState {
    
    @Override
    public void submit(LeaveRequestContext ctx) {
    
        System.out.println("[主管审批] 已提交，等待主管审批");
    }
    
    @Override
    public void approve(LeaveRequestContext ctx) {
    
        System.out.println("[主管审批] 主管同意，进入HRBP审批");
        ctx.setState(new HRBPApprovalState());
    }
    
    @Override
    public void reject(LeaveRequestContext ctx) {
    
        System.out.println("[主管审批] 主管拒绝，流程结束");
        ctx.setState(new EndState());
    }
    
    @Override
    public void cancel(LeaveRequestContext ctx) {
        System.out.println("[主管审批] 已撤销");
        ctx.setState(new EndState());
    }
    
}
