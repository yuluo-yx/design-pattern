package main

import "fmt"

func main() {

	fmt.Println("============== 功能有限的原始类 ================")
	square := &Square{}
	fmt.Println(square.Draw())

	fmt.Println("============== 使用装饰器模式扩展 ================")
	redSquare := NewColorDecorator(square, "red")
	fmt.Println(redSquare.Draw())
}
