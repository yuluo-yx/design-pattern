package main

import "fmt"

// EndState 结束状态

type EndState struct{}

func (s *EndState) Submit(ctx *Context)  { fmt.Println("[结束] 流程已结束，不能再提交") }
func (s *EndState) Approve(ctx *Context) { fmt.Println("[结束] 流程已结束，不能再审批") }
func (s *EndState) Reject(ctx *Context)  { fmt.Println("[结束] 流程已结束，不能再拒绝") }
func (s *EndState) Cancel(ctx *Context)  { fmt.Println("[结束] 流程已结束，不能再撤销") }
