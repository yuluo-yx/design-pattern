package main;

public class ClassAdapter extends LegacyDevice implements JSONDevice {
   
    @Override
    public String getJSONData() {
        
        String xml = getXMLData(); // 调用父类方法

        // 这里应该将 xml 转换为 json，简化处理下
        return "{\"value\":\"" + xml.replaceAll(".*<value>(.*)</value>.*", "$1") + "\"}";
    }
    
}
