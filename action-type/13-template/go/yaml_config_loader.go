package main

import (
	"fmt"
)

type YamlConfigLoader struct {
	BaseConfigLoader
}

func (y *YamlConfigLoader) getReader(filePath string) interface{} {
	fmt.Println("Loading YAML config from:", filePath)
	return nil
}

func (y *YamlConfigLoader) parseConfig(reader interface{}) interface{} {
	fmt.Println("Parsing YAML config...")
	return struct{}{}
}

func (y *YamlConfigLoader) validateConfig(config interface{}) bool {
	fmt.Println("Validating YAML config...")
	return true
}

func (y *YamlConfigLoader) Load(filePath string) interface{} {
	return y.BaseConfigLoader.Load(filePath, y.getReader, y.parseConfig, y.validateConfig)
}
