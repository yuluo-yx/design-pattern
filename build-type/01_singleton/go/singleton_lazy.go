package main

import (
	"sync"
)

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// 单例模式：懒汉式
// 懒汉：在只有用到的时候才去初始化
func GetLazyInstance() *Singleton {

	// once 是 go 标准库提供的函数
	// 常用来初始化配置或者数据库等只执行一次操作
	once.Do(func() {
		lazySingleton = &Singleton{field: "lazy-test"}
	})

	return lazySingleton
}
