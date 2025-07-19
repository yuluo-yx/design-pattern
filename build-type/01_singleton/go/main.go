package main

import (
	"fmt"
)

func main() {

	// 饿汉式
	singleton := GetInstance()
	fmt.Printf("Eager Singleton: %p\n", singleton)

	// 懒汉式
	// 单例模式中获取到的对象始终是同一个对象
	// 即有相同的内存地址
	lazySingleton1 := GetLazyInstance()
	lazySingleton2 := GetLazyInstance()
	fmt.Printf("Lazy Singleton1: %p\n", lazySingleton1)
	fmt.Printf("Lazy Singleton2: %p\n", lazySingleton2)
}
