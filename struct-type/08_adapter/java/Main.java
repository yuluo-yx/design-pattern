package main;

public class Main {
    public static void main(String[] args) {

        // 遗留对象或者需要适配的对象
        LegacyDevice legacy = new LegacyDevice();

        // 对象适配器
        JSONDevice objAdapter = new ObjectAdapter(legacy);
        System.out.println("对象适配器: " + objAdapter.getJSONData());

        // 类适配器
        JSONDevice classAdapter = new ClassAdapter();
        System.out.println("类适配器: " + classAdapter.getJSONData());
    }
}
