package main

import (
	"fmt"
	"strings"
)

type BaseConfigLoader struct{}

func (b *BaseConfigLoader) Load(
	filePath string,
	getReader func(string) interface{},
	parseConfig func(interface{}) interface{},
	validateConfig func(interface{}) bool,
) interface{} {
	if !b.checkConfigFileExists(filePath) {
		panic("Config file does not exist: " + filePath)
	}
	reader := getReader(filePath)
	config := parseConfig(reader)
	if !validateConfig(config) {
		panic("Config validation failed.")
	}
	return config
}

func (b *BaseConfigLoader) checkConfigFileExists(filePath string) bool {
	fmt.Println("Checking if config file exists at:", filePath)
	return true
}

// 工厂方法
func CreateLoader(filePath string) ConfigLoader {
	if strings.HasSuffix(filePath, ".json") {
		return &JSONConfigLoader{}
	} else if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		return &YamlConfigLoader{}
	} else {
		panic("Unsupported config file format: " + filePath)
	}
}
