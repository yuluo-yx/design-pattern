package main

import (
	"fmt"
)

type PersonBuilder struct {

	// 处理构建过程中的错误
	err error

	// 需要建造的对象实例
	person *Person
}

func (pb *PersonBuilder) SetName(name string) *PersonBuilder {

	pb.person.Name = name
	return pb
}

func (pb *PersonBuilder) SetAge(age int8) *PersonBuilder {

	pb.person.Age = age
	return pb
}

func (pb *PersonBuilder) SetAddress(address string) *PersonBuilder {

	pb.person.Address = address
	return pb
}

func (pb *PersonBuilder) Build() *Person {

	// 属性字段检查
	if pb.person.Name == "" {
		pb.err = fmt.Errorf("name is required")
		return nil
	}

	return pb.person
}
