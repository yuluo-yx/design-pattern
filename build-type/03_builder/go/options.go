package main

import (
	"errors"
)

// go 中常用的 options 传递参数方式

var ErrNameRequired = errors.New("name is required")

type DataSourceConfigOption struct {
	Host     string
	Port     int
	Username string
	Password string
}

type DataSourceConfigOptFunc func(option *DataSourceConfigOption)

func NewDataSourceConfig(name string, opts ...DataSourceConfigOptFunc) (*DataSourceCofig, error) {

	// 参数检查
	if name == "" {
		return nil, ErrNameRequired
	}

	option := &DataSourceConfigOption{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "root",
	}

	for _, opt := range opts {
		opt(option)
	}

	// 添加其他的参数检查 ping 测试等

	return &DataSourceCofig{
		Name:     name,
		Host:     option.Host,
		Port:     option.Port,
		Username: option.Username,
		Password: option.Password,
	}, nil

}
