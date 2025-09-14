package main;

import java.io.Reader;


public abstract class AbstractConfigLoader<T> implements ConfigLoader<T> {

    // 简单工厂方法，根据文件后缀返回对应的 ConfigLoader 实例
    public static ConfigLoader<Object> createLoader(String filePath) {
        if (filePath.endsWith(".json")) {
            return new JSONConfigLoader();
        } else if (filePath.endsWith(".yaml") || filePath.endsWith(".yml")) {
            return new YamlConfigLoader();
        } else {
            throw new UnsupportedOperationException("Unsupported config file format: " + filePath);
        }
    }

    // 模板方法
    @Override
    public T load(String filePath) {
        if (!checkConfigFileExists(filePath)) {
            throw new IllegalArgumentException("Config file does not exist: " + filePath);
        }
        Reader reader = getReader(filePath);
        T config = parseConfig(reader);
        if (!validateConfig(config)) {
            throw new IllegalArgumentException("Config validation failed.");
        }
        return config;
    }

    // 步骤1：检查文件
    protected boolean checkConfigFileExists(String filePath) {
        
        System.out.println("Checking if config file exists at: " + filePath);
        return true;
    }

    // 步骤2：获取 Reader，由子类实现
    protected abstract Reader getReader(String filePath);

    // 步骤3：解析配置，由子类实现
    protected abstract T parseConfig(Reader reader);

    // 步骤4：校验配置，可选重写
    protected boolean validateConfig(T config) {

        System.out.println("Validating config...");
        return true;
    }
}
