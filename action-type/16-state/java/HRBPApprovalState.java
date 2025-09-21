package main;

// HRBP 审批状态
public class HRBPApprovalState implements LeaveRequestState {
    
    @Override
    public void submit(LeaveRequestContext ctx) {
    
        System.out.println("[HRBP审批] 已提交，等待HRBP审批");
    }
    
    @Override
    public void approve(LeaveRequestContext ctx) {
    
        System.out.println("[HRBP审批] HRBP同意，流程结束");
        ctx.setState(new EndState());
    }
    @Override
    public void reject(LeaveRequestContext ctx) {
    
        System.out.println("[HRBP审批] HRBP拒绝，流程结束");
        ctx.setState(new EndState());
    }
    @Override
    public void cancel(LeaveRequestContext ctx) {
    
        System.out.println("[HRBP审批] 已撤销");
        ctx.setState(new EndState());
    }
    
}
