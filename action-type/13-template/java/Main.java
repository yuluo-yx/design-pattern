package main;

public class Main {

    public static void main(String[] args) {
    // 直接 new
    ConfigLoader<Object> jsonLoader = new JSONConfigLoader();
    ConfigLoader<Object> yamlLoader = new YamlConfigLoader();

    System.out.println("--- 加载 JSON 配置（new） ---");
    Object jsonConfig = jsonLoader.load("config.json");

    System.out.println("--- 加载 YAML 配置（new） ---");
    Object yamlConfig = yamlLoader.load("config.yaml");

    // 工厂方法
    System.out.println("--- 加载 JSON 配置（工厂） ---");
    ConfigLoader<Object> jsonLoader2 = AbstractConfigLoader.createLoader("config.json");
    Object jsonConfig2 = jsonLoader2.load("config.json");

    System.out.println("--- 加载 YAML 配置（工厂） ---");
    ConfigLoader<Object> yamlLoader2 = AbstractConfigLoader.createLoader("config.yaml");
    Object yamlConfig2 = yamlLoader2.load("config.yaml");
    }
    
}
