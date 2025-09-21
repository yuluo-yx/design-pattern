package main;

// 请假单状态接口
public interface LeaveRequestState {
    
    void submit(LeaveRequestContext ctx);
    void approve(LeaveRequestContext ctx);
    void reject(LeaveRequestContext ctx);
    void cancel(LeaveRequestContext ctx);
}
