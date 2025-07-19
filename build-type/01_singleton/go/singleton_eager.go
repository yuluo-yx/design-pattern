package main

// 实例
var singleton *Singleton

// 单例模式 - 饿汉式
func init() {
	singleton = &Singleton{field: "test"}
}

// 获取实例
func GetInstance() *Singleton {

	return singleton
}
