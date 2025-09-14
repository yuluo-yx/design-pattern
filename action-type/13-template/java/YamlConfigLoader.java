
package main;

import java.io.Reader;

public class YamlConfigLoader extends AbstractConfigLoader<Object> {

    @Override
    protected Reader getReader(String filePath) {
        System.out.println("Loading YAML config from: " + filePath);
        // 实际实现应返回文件读取器
        return null;
    }

    @Override
    protected Object parseConfig(Reader reader) {
        System.out.println("Parsing YAML config...");
        // 实际实现应解析 YAML 并返回配置对象
        return new Object();
    }
}
