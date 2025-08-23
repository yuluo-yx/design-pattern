package main;

// 对象适配器
public class ObjectAdapter implements JSONDevice {

    // 持有遗留类对象
    private LegacyDevice legacyDevice;

    public ObjectAdapter(LegacyDevice legacyDevice) {
        
        this.legacyDevice = legacyDevice;
    }

    @Override
    public String getJSONData() {

        // 简化处理，实际可用 XML/JSON 库转换
        return "{\"value\":\"Hello from legacy device\"}";
    }

}