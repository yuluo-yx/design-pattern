package main;

// 消息发送者接口
interface IMsgSender {
    void send(String msg);
}

// 邮件消息发送者实现
class EmailMsgSender implements IMsgSender {

    private String[] email;

    public EmailMsgSender(String[] email) {
        this.email = email;
    }

    @Override
    public void send(String msg) {

        for (String e : email) {
            System.out.println("Sending email to " + e + ": " + msg);
        }
    }
}

class DingTalkMsgSender implements IMsgSender {

    private String dingtalk;

    public DingTalkMsgSender(String dingtalk) {
        this.dingtalk = dingtalk;
    }

    @Override
    public void send(String msg) {

        System.out.println("Sending DingTalk message to " + this.dingtalk + ": " + msg);
    }
}

interface INotify {
    void notify(String msg);
}

// Error 通知实现
class ErrorNotify implements INotify {

    // 桥接模式组合接口
    private IMsgSender msgSender;

    public ErrorNotify(IMsgSender msgSender) {
        this.msgSender = msgSender;
    }

    @Override
    public void notify(String msg) {

        // 调用消息发送者的发送方法
        this.msgSender.send(msg);
    }
}

// Warn 类型的 notify 实现
class WarnNotify implements INotify {

    // 桥接模式组合接口
    private IMsgSender msgSender;

    public WarnNotify(IMsgSender msgSender) {
        this.msgSender = msgSender;
    }

    @Override
    public void notify(String msg) {

        // 调用消息发送者的发送方法
        this.msgSender.send(msg);
    }
}

// 没有用处，仅仅只是为了标记名字。
public class Bridge {
}
