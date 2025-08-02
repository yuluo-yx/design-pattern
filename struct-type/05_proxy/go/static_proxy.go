package main

import (
	"errors"
	"fmt"
	"time"
)

var StaticProxyError = errors.New("static proxy has error")

// 静态工厂接口定义
type IStaticProxy interface {
	CheckAccess() (bool, error)
	PreHandler() error
	PostHandler() error
}

// 静态代理类，代理访问原始的 userService 类服务
type StaticProxy struct {
	// 被代理的原始类
	// 代理的对象属于私有变量，只能内部访问，不应被外部调用者看到
	userService IUserService

	// 辅助字段：统计方法的执行时间
	startTime time.Time
}

func NewStaticProxy(userService IUserService) *StaticProxy {

	return &StaticProxy{
		userService: userService,
	}
}

func (sp *StaticProxy) GetUser(name string) (user *User, err error) {

	// 执行前置处理
	_ = sp.PreHandler()
	// 执行后置处理
	defer func() {
		_ = sp.PostHandler()
	}()

	// 检查是否需要调用原始类
	if ok, err := sp.CheckAccess(); !ok || err != nil {
		return nil, StaticProxyError
	}

	// 调用真实服务
	return sp.userService.GetUser(name)
}

// 前置处理，加入开始时间统计方法执行耗时等
func (sp *StaticProxy) PreHandler() error {
	sp.startTime = time.Now()
	fmt.Println("Static proxy pre handler....")
	return nil
}

func (sp *StaticProxy) PostHandler() error {

	// 打印方法执行耗时
	costSec := time.Since(sp.startTime).Seconds()
	fmt.Printf("Static proxy post handler... 耗时 %.7f 秒\n", costSec)
	return nil
}

func (sp *StaticProxy) CheckAccess() (bool, error) {

	// 检查是否需要调用原始类
	return true, nil
}
