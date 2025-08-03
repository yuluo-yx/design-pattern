package main

import (
	"fmt"
)

// IMsgSender 消息发送者接口
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender 邮件消息发送者
// 短信，钉钉，电话等多种实现
type EmailMsgSender struct {
	email []string
}

func NewEmailMsgSender(email []string) *EmailMsgSender {

	return &EmailMsgSender{email: email}
}

// 邮件消息发送，实现消息发送者接口
func (e *EmailMsgSender) Send(msg string) error {

	// 发送邮件逻辑
	for _, to := range e.email {
		// 这里模拟发送邮件
		fmt.Println("Sending email to:", to, "with message:", msg)
	}
	return nil
}

type DingTalkMsgSender struct {
	// 钉钉机器人地址
	webhook string
}

func NewDingTalkMsgSender(webhook string) *DingTalkMsgSender {

	return &DingTalkMsgSender{webhook: webhook}
}

func (d *DingTalkMsgSender) Send(msg string) error {

	// 发送钉钉消息逻辑
	// 这里模拟发送钉钉消息
	fmt.Println("Sending DingTalk message to:", d.webhook, "with message:", msg)
	return nil
}

// INotify 消息通知接口
type INotify interface {
	Notify(msg string) error
}

// ErrNotify 错误消息通知
type ErrNotify struct {
	// 消息通过不同的发送方式发送出去
	// 使用消息发送接口
	msgSender IMsgSender
}

// 桥接模式 组合接口
// NewErrNotify 将消息发送方式接口组合到消息发送接口中
// 通过参数传入，使用什么方式发送由接口实现决定，以此来解耦两种不同类型
func NewErrNotify(msgSender IMsgSender) *ErrNotify {

	return &ErrNotify{msgSender: msgSender}
}

func (n *ErrNotify) Notify(msg string) error {

	// 发送消息通知
	return n.msgSender.Send(msg)
}
