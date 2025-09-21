package main

import "fmt"

// HRBPApprovalState HRBP审批

type HRBPApprovalState struct{}

func (s *HRBPApprovalState) Submit(ctx *Context) {
	fmt.Println("[HRBP审批] 已提交，等待HRBP审批")
}

func (s *HRBPApprovalState) Approve(ctx *Context) {
	fmt.Println("[HRBP审批] HRBP同意，流程结束")
	ctx.SetState(&EndState{})
}

func (s *HRBPApprovalState) Reject(ctx *Context) {
	fmt.Println("[HRBP审批] HRBP拒绝，流程结束")
	ctx.SetState(&EndState{})
}

func (s *HRBPApprovalState) Cancel(ctx *Context) {
	fmt.Println("[HRBP审批] 已撤销")
	ctx.SetState(&EndState{})
}
