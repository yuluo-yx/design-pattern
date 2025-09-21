package main;

// 草稿状态：可提交、可撤销
public class DraftState implements LeaveRequestState {

    @Override
    public void submit(LeaveRequestContext ctx) {
    
        System.out.println("[草稿] 提交请假单，进入主管审批");
        ctx.setState(new ManagerApprovalState());
    }
    
    @Override
    public void approve(LeaveRequestContext ctx) {
    
        System.out.println("[草稿] 不能直接审批");
    }
    
    @Override
    public void reject(LeaveRequestContext ctx) {
    
        System.out.println("[草稿] 不能拒绝");
    }
    
    @Override
    public void cancel(LeaveRequestContext ctx) {
    
        System.out.println("[草稿] 已撤销");
        ctx.setState(new EndState());
    }
}
