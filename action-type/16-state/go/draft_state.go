package main

import "fmt"

// DraftState 草稿状态

type DraftState struct{}

func (s *DraftState) Submit(ctx *Context) {

	fmt.Println("[草稿] 提交请假单，进入主管审批")
	ctx.SetState(&ManagerApprovalState{})
}

func (s *DraftState) Approve(ctx *Context) {

	fmt.Println("[草稿] 不能直接审批")
}

func (s *DraftState) Reject(ctx *Context) {

	fmt.Println("[草稿] 不能拒绝")
}

func (s *DraftState) Cancel(ctx *Context) {

	fmt.Println("[草稿] 已撤销")
	ctx.SetState(&EndState{})
}
