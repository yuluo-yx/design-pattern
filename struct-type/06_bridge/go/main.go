package main

import (
	"fmt"
)

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
