package main

func main() {

	// 创建 DOM 树
	div := NewElementNode("div")
	div.SetAttribute("class", "container")

	p1 := NewElementNode("p")
	p1.AppendChild(NewTextNode("这是第一个段落"))

	p2 := NewElementNode("p")
	p2.SetAttribute("style", "color: blue")
	p2.AppendChild(NewTextNode("这是第二个段落"))

	div.AppendChild(p1)
	div.AppendChild(p2)

	// 渲染整个 DOM 树
	div.Render()
}
