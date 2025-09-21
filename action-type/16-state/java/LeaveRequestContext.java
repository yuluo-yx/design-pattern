package main;

// LeaveRequestContext 持有当前状态和业务数据
public class LeaveRequestContext {
    
    private LeaveRequestState state;
    
    private String applicant;
    
    private String reason;

    public LeaveRequestContext(String applicant, String reason) {
        this.applicant = applicant;
        this.reason = reason;
        this.state = new DraftState();
    }

    public void setState(LeaveRequestState state) {
        this.state = state;
    }

    public LeaveRequestState getState() {
        return state;
    }

    public String getApplicant() {
        return applicant;
    }

    public String getReason() {
        return reason;
    }

    public void submit() {
        state.submit(this);
    }
    
    public void approve() {
        state.approve(this);
    }
    
    public void reject() {
        state.reject(this);
    }
    
    public void cancel() {
        state.cancel(this);
    }

}
