package main

import (
	"fmt"
	"reflect"
)

// 基于 go 反射实现动态代理

// DynamicReflectProxyHandler 代理类
type DynamicReflectProxy struct {
	target any
}

// PreHandler 前置处理
func (rh *DynamicReflectProxy) PreHandler(method string, args []any) {

	fmt.Printf("Before calling %s\n", method)
}

// PostHandler 后置处理
func (rh *DynamicReflectProxy) PostHandler(method string, args []any) {

	fmt.Printf("After calling %s\n", method)
}

// GetUser 代理方法
// 直接实现 userService 的 GetUser 方法
// 也可以不实现，直接使用反射调用
func (rh *DynamicReflectProxy) GetUser(name string) (*User, error) {

	// 获取方法信息
	methodName := "GetUser"
	args := []any{name}

	// 前置处理
	rh.PreHandler(methodName, args)
	defer func() {
		// 后置处理
		rh.PostHandler(methodName, args)
	}()

	// 使用反射调用目标方法
	targetValue := reflect.ValueOf(rh.target)
	results := targetValue.MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(name)})

	// 处理返回值
	user := results[0].Interface().(*User)
	var err error
	if !results[1].IsNil() {
		err = results[1].Interface().(error)
	}

	return user, err
}

// NewProxy 创建代理
func NewDynamicReflectProxy(target any) IUserService {

	return &DynamicReflectProxy{target: target}
}
