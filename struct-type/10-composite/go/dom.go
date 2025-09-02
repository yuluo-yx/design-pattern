package main

import (
	"fmt"
)

// DOMNode 定义所有 DOM 节点的通用行为
type DOMNode interface {
	Render()
}

// BaseNode 节点的基本属性
type BaseNode struct {
	tagName string
	content string
}

// TextNode 文本节点
type TextNode struct {
	BaseNode
}

// NewTextNode 创建文本节点
func NewTextNode(content string) *TextNode {

	return &TextNode{
		BaseNode: BaseNode{
			tagName: "#text",
			content: content,
		},
	}
}

// Render 渲染文本节点
func (t *TextNode) Render() {
	fmt.Print(t.content)
}

// ElementNode 元素节点
type ElementNode struct {
	BaseNode
	children   []DOMNode
	attributes map[string]string
}

// NewElementNode 创建元素节点
func NewElementNode(tagName string) *ElementNode {

	return &ElementNode{
		BaseNode: BaseNode{
			tagName: tagName,
		},
		children:   make([]DOMNode, 0),
		attributes: make(map[string]string),
	}
}

// SetAttribute 设置属性
func (e *ElementNode) SetAttribute(name, value string) {
	e.attributes[name] = value
}

// AppendChild 添加子节点
func (e *ElementNode) AppendChild(child DOMNode) {
	e.children = append(e.children, child)
}

// Render 渲染元素节点
func (e *ElementNode) Render() {
	// 渲染开始标签和属性
	fmt.Print("<", e.tagName)
	for name, value := range e.attributes {
		fmt.Printf(" %s=\"%s\"", name, value)
	}
	fmt.Print(">")

	// 渲染子节点
	for _, child := range e.children {
		child.Render()
	}

	// 渲染结束标签
	fmt.Print("</", e.tagName, ">")
}
