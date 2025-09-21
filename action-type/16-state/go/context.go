package main

// Context 持有当前状态和业务数据

type Context struct {
	State     State
	Applicant string
	Reason    string
}

func NewContext(applicant, reason string) *Context {
	return &Context{
		State:     &DraftState{},
		Applicant: applicant,
		Reason:    reason,
	}
}

func (c *Context) SetState(state State) {
	c.State = state
}

func (c *Context) Submit()  { c.State.Submit(c) }
func (c *Context) Approve() { c.State.Approve(c) }
func (c *Context) Reject()  { c.State.Reject(c) }
func (c *Context) Cancel()  { c.State.Cancel(c) }
