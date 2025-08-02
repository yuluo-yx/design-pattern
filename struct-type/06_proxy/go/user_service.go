package main

import "fmt"

// 初始化一个 map 用来替代数据库服务
var userDB map[string]*User = make(map[string]*User, 2)

func init() {
	userDB["Golang"] = &User{Name: "Golang", Age: 18}
	userDB["Java"] = &User{Name: "Java", Age: 18}
}

type User struct {
	Name string
	Age  int8
}

type IUserService interface {
	GetUser(name string) (*User, error)
}

type UserService struct {
}

// 创建一个 userService 实例类
func NewUserService() IUserService {

	return &UserService{}
}

func (u *UserService) GetUser(name string) (user *User, err error) {

	fmt.Println("Get user info, key is:", name)
	return userDB[name], nil
}
