# 建造者模式 （Builder Design Pattern）

建造者模式的主要目的是将一个复杂对象的构建过程与其表示相分离，从而可以创建具有不同表示形式的对象。

## 1. 应用

1. Java 中的 StringBuilder
2. Spring AI Alibaba 框架中的大模型 Options 参数

### 1.1 优点

1. 适用于具有很多复杂属性的类；
2. 将复杂对象的创建和使用分离

### 1.2 缺点

1. 增加类数量；
2. 内部修改，如果类属性发生变化，建造者也要更改。

## 2. 实现分析

### 2.1 Java 

```java
public class Person {

    // 实例类属性
    private String name;

    private Integer age;

    private String address;

    public String getName() {
        return name;
    }

    public Integer getAge() {
        return age;
    }

    public String getAddress() {
        return address;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setAge(Integer age) {
        this.age = age;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    // 实例类构造方法
    public Person(String name, Integer age, String address) {
        this.name = name;
        this.age = age;
        this.address = address;
    }

    // 建造者模式的常用方式是将实例类和建造者类定义在同一个类文件中，
    // 并提供对外可以获取的 builder 方法用于构建对象。
    public static Builder builder() {

        return new Builder();
    }

    public static class Builder {

        private String name;
        private Integer age;
        private String address;

        // 无参构造私有化
        // 只允许使用 builder 方法来创建 Builder 对象
  			private Builder() {
        }

        public Builder name(String name) {
            this.name = name;
            return this;
        }

        public Builder age(int age) {
            this.age = age;
            return this;
        }

        public Builder address(String address) {
            this.address = address;
            return this;
        }

        // 构建方法，返回 Person 对象
        public Person build() {

            // 验证必填字段
            if (this.name == null || this.name.isEmpty()) {
                throw new IllegalArgumentException("Name cannot be null or empty");
            }

            return new Person(name, age, address);
        }

    }

}
```

### 2.2 Go

在 Go 中，通常不直接使用类似 Java 的 Builder 用法，常见的做法是必填参数直接传递，可选参数通过传递可变的方法进行创建。

#### 普通 Builder 模式

```go
package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int8
	Address string
}

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
```

#### Option 用法

```go
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
```

## 3. 相关面试题

Q: 建造者与工厂模式的区别？
A: 注重点不同: 建造者模式 更注重于方法的调用顺序 ; 工厂模式注重于创建实例产品 , 不关心方法调用的顺序 ;
创建对象力度不同: 创建对象的力度不同 , 建造者模式可以创建复杂的产品 , 由各种复杂的内部属性组成, 工厂模式创建出来的都是相同的实例对象。

```java
var options = DashScopeChatOptions.builder()
        .withEnableSearch(true)
        .withModel(DashScopeApi.ChatModel.DEEPSEEK_V3.getValue())
        .withSearchOptions(DashScopeApi.SearchOptions.builder()
                .forcedSearch(true)
            .enableSource(true)
                .searchStrategy("pro")
                .enableCitation(true)
                .citationFormat("[<number>]")
                .build()
        ).withTemperature(0.7)
        .build();
```
