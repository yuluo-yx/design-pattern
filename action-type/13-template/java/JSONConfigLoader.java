
package main;

import java.io.Reader;

public class JSONConfigLoader extends AbstractConfigLoader<Object> {

    @Override
    protected Reader getReader(String filePath) {
        System.out.println("Loading JSON config from: " + filePath);
        // 实际实现应返回文件读取器
        return null;
    }

    @Override
    protected Object parseConfig(Reader reader) {
        System.out.println("Parsing JSON config...");
        // 实际实现应解析 JSON 并返回配置对象
        return new Object();
    }
}
