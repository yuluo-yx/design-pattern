package main

import (
	"fmt"
)

type JSONConfigLoader struct {
	BaseConfigLoader
}

func (j *JSONConfigLoader) getReader(filePath string) interface{} {
	fmt.Println("Loading JSON config from:", filePath)
	return nil
}

func (j *JSONConfigLoader) parseConfig(reader interface{}) interface{} {
	fmt.Println("Parsing JSON config...")
	return struct{}{}
}

func (j *JSONConfigLoader) validateConfig(config interface{}) bool {
	fmt.Println("Validating JSON config...")
	return true
}

func (j *JSONConfigLoader) Load(filePath string) interface{} {
	return j.BaseConfigLoader.Load(filePath, j.getReader, j.parseConfig, j.validateConfig)
}
