package main

import "fmt"

// ManagerApprovalState 主管审批

type ManagerApprovalState struct{}

func (s *ManagerApprovalState) Submit(ctx *Context) {
	fmt.Println("[主管审批] 已提交，等待主管审批")
}

func (s *ManagerApprovalState) Approve(ctx *Context) {
	fmt.Println("[主管审批] 主管同意，进入HRBP审批")
	ctx.SetState(&HRBPApprovalState{})
}

func (s *ManagerApprovalState) Reject(ctx *Context) {
	fmt.Println("[主管审批] 主管拒绝，流程结束")
	ctx.SetState(&EndState{})
}

func (s *ManagerApprovalState) Cancel(ctx *Context) {
	fmt.Println("[主管审批] 已撤销")
	ctx.SetState(&EndState{})
}
