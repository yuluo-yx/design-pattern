package main

import "fmt"

func main() {

	ctx := NewContext("张三", "家中有事请假2天")
	fmt.Println("--- 发起请假 ---")
	ctx.Submit()
	fmt.Println("--- 主管审批同意 ---")
	ctx.Approve()
	fmt.Println("--- HRBP 审批同意 ---")
	ctx.Approve()
	fmt.Println("--- 尝试再次提交 ---")
	ctx.Submit()

	fmt.Println("--- 新流程：发起请假，主管拒绝 ---")
	ctx2 := NewContext("李四", "身体不适请假1天")
	ctx2.Submit()
	ctx2.Reject()

	fmt.Println("--- 新流程：发起请假，撤销 ---")
	ctx3 := NewContext("王五", "家中有事请假3天")
	ctx3.Cancel()
}
