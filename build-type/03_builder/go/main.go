package main

import (
	"fmt"
)

func main() {

	// 创建一个 PersonBuilder 实例
	builder := &PersonBuilder{
		person: &Person{},
	}

	// 使用链式调用设置属性
	person1 := builder.SetName("John").
		SetAge(30).
		SetAddress("123 Main St").
		Build()
	// 检查构建过程中是否有错误
	if builder.err != nil {
		fmt.Println("Error building person:", builder.err)
		return
	}

	person2 := builder.SetName("Jane").
		SetAge(25).
		SetAddress("456 Elm St").
		Build()
	// 检查构建过程中是否有错误
	if builder.err != nil {
		fmt.Println("Error building person:", builder.err)
		return
	}

	// 输出构建的 Person 对象
	fmt.Printf("person1: %p\n", person1)
	fmt.Printf("person2: %p\n", person2)

	// options 用法
	dataSourceConfig1, err := NewDataSourceConfig("mydb1",
		func(option *DataSourceConfigOption) {
			option.Host = "localhost"
			option.Port = 3306
		},
	)
	if err != nil {
		fmt.Println("Error creating data source config:", err)
		return
	}

	dataSourceConfig2, err := NewDataSourceConfig("mydb2",
		func(option *DataSourceConfigOption) {
			option.Host = "localhost"
			option.Port = 3306
			option.Username = "admin"
			option.Password = "admin123"
		},
	)
	if err != nil {
		fmt.Println("Error creating data source config:", err)
		return
	}

	fmt.Printf("dataSourceConfig1: %p, Name: %s, Host: %s, Port: %d, Username: %s, Password: %s\n", dataSourceConfig1, dataSourceConfig1.Name, dataSourceConfig1.Host, dataSourceConfig1.Port, dataSourceConfig1.Username, dataSourceConfig1.Password)
	fmt.Printf("dataSourceConfig2: %p, Name: %s, Host: %s, Port: %d, Username: %s, Password: %s\n", dataSourceConfig2, dataSourceConfig2.Name, dataSourceConfig2.Host, dataSourceConfig2.Port, dataSourceConfig2.Username, dataSourceConfig2.Password)

}
