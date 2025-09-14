package main

import "fmt"

func main() {
	fmt.Println("--- 加载 JSON 配置（new） ---")
	jsonLoader := &JSONConfigLoader{}
	jsonLoader.Load("config.json")

	fmt.Println("--- 加载 YAML 配置（new） ---")
	yamlLoader := &YamlConfigLoader{}
	yamlLoader.Load("config.yaml")

	fmt.Println("--- 加载 JSON 配置（工厂） ---")
	jsonLoader2 := CreateLoader("config.json")
	jsonLoader2.Load("config.json")

	fmt.Println("--- 加载 YAML 配置（工厂） ---")
	yamlLoader2 := CreateLoader("config.yaml")
	yamlLoader2.Load("config.yaml")
}
