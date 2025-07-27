package main;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.io.Serial;
import java.io.Serializable;

public class ChatModel implements Serializable, Cloneable {

    @Serial
    private static final long serialVersionUID = 1L;

    private String apiKey;

    private String apiSecret;

    private ChatOption chatOption;

    public ChatModel() {
    }

    public ChatModel(String apiKey, String apiSecret, ChatOption chatOption) {
        this.apiKey = apiKey;
        this.apiSecret = apiSecret;
        this.chatOption = chatOption;
    }

    public String getApiKey() {
        return apiKey;
    }

    public void setApiKey(String apiKey) {
        this.apiKey = apiKey;
    }

    public String getApiSecret() {
        return apiSecret;
    }

    public void setApiSecret(String apiSecret) {
        this.apiSecret = apiSecret;
    }

    public ChatOption getChatOption() {
        return chatOption;
    }

    public void setChatOption(ChatOption chatOption) {
        this.chatOption = chatOption;
    }

    // Object 类中的 clone 实现，不做序列化
    @Override
    public ChatModel clone() {
        try {
            return (ChatModel) super.clone();
        } catch (CloneNotSupportedException e) {
            throw new AssertionError();
        }
    }

    // 深拷贝方法
    // 序列化实现
    public ChatModel deepClone1() {

        var chatModel = new ChatModel();

        try {
            // 序列化
            ByteArrayOutputStream bos = new ByteArrayOutputStream();
            ObjectOutputStream oos = new ObjectOutputStream(bos);
            oos.writeObject(this);
            oos.flush();
            oos.close();

            // 反序列化
            ByteArrayInputStream bis = new ByteArrayInputStream(bos.toByteArray());
            ObjectInputStream ois = new ObjectInputStream(bis);
            chatModel = (ChatModel) ois.readObject();
            return chatModel;
        } catch (IOException | ClassNotFoundException e) {
            throw new RuntimeException("Clone failed", e);
        }
    }

    // 手动处理引用
    public ChatModel deepClone2() {

        var chatModel = new ChatModel();

        // 处理基本类型
        chatModel.setApiKey(this.apiKey);
        chatModel.setApiSecret(this.apiSecret);

        // 处理引用类型
        if (this.chatOption != null) {
            chatModel.setChatOption(this.chatOption.clone());
        }
        return chatModel;
    }

}

class ChatOption implements Serializable, Cloneable {

    @Serial
    private static final long serialVersionUID = 1L;

    private String model;

    public ChatOption() {
    }

    public ChatOption(String model) {
        this.model = model;
    }

    @Override
    public ChatOption clone() {
        try {
            return (ChatOption) super.clone();
        } catch (CloneNotSupportedException e) {
            throw new AssertionError();
        }
    }

}
