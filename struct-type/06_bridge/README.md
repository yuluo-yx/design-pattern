# 桥接模式 Bridge Design Pattern）

桥接模式：将抽象部分与它的实现部分分离，使得可以独立变化。又称为柄体(Handle and Body)模式或接口(Interface)模式。其涉及一个接口，它充当一个桥，使得具体类可以在不影响客户端代码的情况下改变。

桥接模式将继承关系转化为组合关系，从而减少类之间的紧密耦合度，使得系统更加灵活和可扩展。

## 1. 应用

桥接模式的定义非常模糊不清晰，在实际应用中较少（当然大多数情况下可能是用了，但是并不知道这是桥接模式）。本身是试图将一个庞大且复杂的类拆分成更细粒度的类，继而通过某种更合理的结构进行组装。

在 Spring 中的 JdbcTemplate 有应用。

### 1.1 优点

1. 分离抽象接口及其实现部分；
2. 桥接模式有时类似于多继承方案，但是多继承方案违背了类的单一职责原则（即一个类只有一个变化的原因），复用性比较差，而且多继承结构中类的个数非常庞大，桥接模式是比多继承方案更好的解决方法；
3. 桥接模式提高了系统的可扩充性，在两个变化维度中任意扩展一个维度，都不需要修改原有代码。

### 1.2 缺点

1. 桥接模式的引入会增加系统的理解与设计难度，由于聚合关联关系建立在抽象层，要求对抽象编程；
2. 桥接模式要求正确识别出系统中两个独立变化的维度，因此其使用范围具有一定的局限性。

## 2. 实现分析

以 Go 为例：

```golang
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
```

main 客户端调用：

```golang
func main() {

	// 邮件发送
	eSender := NewEmailMsgSender([]string{"123@test.com", "456@test.com"})
	// 直接调用消息发送方法发送消息
	eSender.Send("This is a test email message.")

	fmt.Println("--------------------------------------------------------------------------------------")

	// 组合消息通知接口
	// 消息通知方式可以在不变更消息发送者的情况下进行扩展
	// 例如：如果需要发送错误告警通知，可以使用 ErrNotify 结构体来包装消息发送者
	emailErrNotify := NewErrNotify(eSender)
	emailErrNotify.Notify("This is a test error notification message.")

	fmt.Println("--------------------------------------------------------------------------------------")

	// 钉钉发送
	dSender := NewDingTalkMsgSender("https://dingtalk-test.com")
	dSender.Send("This is a test DingTalk message.")

	fmt.Println("--------------------------------------------------------------------------------------")

	dingTalkErrNotify := NewErrNotify(dSender)
	dingTalkErrNotify.Notify("This is a test error notification message.")

}
```

## 3. 参考文档

1. https://design-patterns.readthedocs.io/zh-cn/latest/structural_patterns/bridge.html
2. https://lailin.xyz/post/bridge.html
