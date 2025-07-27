package main;

public class Main {
  
    public static void main(String[] args) {

        ChatModel oldChatModel = new ChatModel("http://127.0.0.1", "sk-abcedefg", new ChatOption("gpt-3.5-turbo"));
        System.out.println("浅拷贝，对 clone 方法不做任何修改 ==============================");
        ChatModel newClone1 = oldChatModel.clone();
        System.out.println("Old ChatModel: " + oldChatModel + "引用对象地址: " + oldChatModel.getChatOption());
        System.out.println("New Clone 1 : " + newClone1 + "引用对象地址: " + newClone1.getChatOption());

        System.out.println("深拷贝 使用 jdk 序列化实现 ==============================");
        ChatModel newClone2 = oldChatModel.deepClone1();
        System.out.println("Old ChatModel: " + oldChatModel + "引用对象地址: " + oldChatModel.getChatOption());
        System.out.println("New Clone 2 : " + newClone2 + "引用对象地址: " + newClone2.getChatOption());


        System.out.println("深拷贝 手动处理引用类型实现 ==============================");
        ChatModel newClone3 = oldChatModel.deepClone1();
        System.out.println("Old ChatModel: " + oldChatModel + "引用对象地址: " + oldChatModel.getChatOption());
        System.out.println("New Clone 2 : " + newClone3 + "引用对象地址: " + newClone3.getChatOption());
    }

}
