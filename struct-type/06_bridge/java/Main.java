package main;

public class Main {

    public static void main(String[] args) {

        EmailMsgSender emailMsgSender = new EmailMsgSender(new String[]{"a@a.com", "b@b.com"});
        // 调用邮件发送实现的发送方法
        emailMsgSender.send("Hello, this is an email message!");

        System.out.println("-------------------------------------------------------------------");
        // 使用邮件发送告警信息
        ErrorNotify errorNotify = new ErrorNotify(emailMsgSender);
        errorNotify.notify("Error: Server down.");

        System.out.println("-------------------------------------------------------------------");

        DingTalkMsgSender dingtalk = new DingTalkMsgSender("dingtalk");
        dingtalk.send("Hello, this is a DingTalk message!");
        System.out.println("-------------------------------------------------------------------");

        // 使用钉钉发送告警信息
        WarnNotify warnNotify = new WarnNotify(dingtalk);
        warnNotify.notify("Warning: High CPU usage detected.");
    }

}
