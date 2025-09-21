package main

// State 接口，定义请假流程的所有操作
// 每个状态都要实现这些方法

type State interface {
	Submit(ctx *Context)
	Approve(ctx *Context)
	Reject(ctx *Context)
	Cancel(ctx *Context)
}
