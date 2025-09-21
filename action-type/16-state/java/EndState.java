package main;

// 结束状态
public class EndState implements LeaveRequestState {
    
    @Override
    public void submit(LeaveRequestContext ctx) {
    
        System.out.println("[结束] 流程已结束，不能再提交");
    }
    
    @Override
    public void approve(LeaveRequestContext ctx) {
    
        System.out.println("[结束] 流程已结束，不能再审批");
    }
    
    @Override
    public void reject(LeaveRequestContext ctx) {
    
        System.out.println("[结束] 流程已结束，不能再拒绝");
    }
    
    @Override
    public void cancel(LeaveRequestContext ctx) {
    
        System.out.println("[结束] 流程已结束，不能再撤销");
    }
    
}
