package main

import (
	"fmt"
	"os"
)

func main() {

	// 生成动态代理文件
	// 检查是否是代码生成模式
	if len(os.Args) > 1 && os.Args[1] == "generate" {
		if err := GenerateProxy(); err != nil {
			fmt.Printf("Error generating proxy: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// 需要被代理的原始类对象
	realUserService := NewUserService()

	fmt.Println("=============== 静态代理 ===============")
	staticProxy := NewStaticProxy(realUserService)
	user, err := staticProxy.GetUser("Golang")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("user: %+v. \n", user)
	fmt.Printf("=============== 静态代理 =============== \n")

	fmt.Println("=============== 动态代理 反射 ===============")
	dynamicReflectProxy := NewDynamicReflectProxy(realUserService)
	drUser, _ := dynamicReflectProxy.GetUser("Golang")
	fmt.Printf("Dynamic Proxy Reflect Got user: %+v\n", drUser)
	fmt.Printf("=============== 动态代理 反射 ============== \n")

	// 如果需要执行 go generate 的代理方式，需要按照以下步骤进行
	// 1. 执行 make go-gen(go generate)
	// 2. 取消下面注释
	// 3. 执行 make go-gen-run
	// fmt.Println("=============== 动态代理 go generate ==============")
	// proxy := NewDynamicGenerateProxy(realUserService)
	// dGenUser, err := proxy.GetUser("Golang")
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Got user: %+v\n", dGenUser)
	// fmt.Println("=============== 动态代理 go generate ==============")

}
