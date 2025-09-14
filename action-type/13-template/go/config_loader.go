package main

type ConfigLoader interface {
	Load(filePath string) interface{}
}
