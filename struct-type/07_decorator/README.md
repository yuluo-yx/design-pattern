# 装饰器模式 （Decorator Design Pattern）

装饰器模式主要对现有的类进行包装，在不改变原有类对象和定义的情况下，扩展对象的其他功能。是结构型设计模式。例如 Spring 框架中的各种 XXXWrapper 或者 XXXDecotator 类。

## 1. 应用

Spring 中的 AOP 设计模式可以理解为装饰器模式的一种实现，在不改变原有类的前提下插入某段日志代码或者检查等。

Spring MVC 中的 HttpHeadResponseDecorator 类等

> 注意区别代理模式：设计模式的理念是互通的。

### 1.1 优点

1. 相比较于继承类，实现方式更加灵活，是继承的一个替代方式；
2. 装饰类和原始类互相独立，没有较高耦合。

### 1.2 缺点

1. 类比继承，结构并不清晰；
2. 包括层数太多时，系统复杂度和理解成本增加。

## 2. 实现分析

以 Golang 代码为例进行分析：

```go
package main

type IDraw interface {
	Draw() string
}

// 原始的正方形绘制方法实现
type Square struct{}

func (s *Square) Draw() string {

	return "Drawing a square"
}

// 使用装饰器模式扩展
type ColorDecorator struct {
	drawable IDraw
	color    string
}

func NewColorDecorator(drawable IDraw, color string) *ColorDecorator {

	return &ColorDecorator{

		// 把原始类当作包装类的变量注入
		drawable: drawable,
		color:    color,
	}
}

func (c *ColorDecorator) Draw() string {

	return c.drawable.Draw() + " with color " + c.color
}
```
