package main

import (
	"fmt"
)

// Subscriber 订阅者接口
type Subscriber interface {
	Receive(news string)
}

// NewspaperOffice 报社接口
type NewspaperOffice interface {
	Subscribe(sub Subscriber)
	Unsubscribe(sub Subscriber)
	Publish(news string)
}

// NewsAgency 具体报社
type NewsAgency struct {
	subscribers []Subscriber
}

func (n *NewsAgency) Subscribe(sub Subscriber) {

	n.subscribers = append(n.subscribers, sub)
}

func (n *NewsAgency) Unsubscribe(sub Subscriber) {

	for i, s := range n.subscribers {
		if s == sub {
			n.subscribers = append(n.subscribers[:i], n.subscribers[i+1:]...)
			break
		}
	}
}

func (n *NewsAgency) Publish(news string) {

	fmt.Printf("报社发布新闻: %s\n", news)

	for _, sub := range n.subscribers {
		sub.Receive(news)
	}
}

// Reader 具体订阅者
type Reader struct {
	name string
}

func (r *Reader) Receive(news string) {

	fmt.Printf("%s 收到报纸: %s\n", r.name, news)
}

func main() {

	agency := &NewsAgency{}

	readerA := &Reader{name: "yuluo"}
	readerB := &Reader{name: "musheng"}

	agency.Subscribe(readerA)
	agency.Subscribe(readerB)

	agency.Publish("2025年9月6日头条新闻：Go 观察者模式示例！")

	agency.Unsubscribe(readerA)
	agency.Publish("2025年9月7日头条新闻：只剩下 musheng 在订阅！")
}
