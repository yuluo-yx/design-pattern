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
