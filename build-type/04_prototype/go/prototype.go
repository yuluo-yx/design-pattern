package main

import (
	"encoding/json"
	"time"
)

// 使用对象序列化来完成原型模式的深拷贝

type KeyWord struct {
	Word     string
	Visit    int
	UpdateAt *time.Time
}

// keyword 的 Clone 函数实现
func (k *KeyWord) Clone() *KeyWord {

	var newKeyWords KeyWord

	// 序列化
	b, err := json.Marshal(k)
	if err != nil {
		panic(err)
	}

	// 反序列化
	err = json.Unmarshal(b, &newKeyWords)
	if err != nil {
		panic(err)
	}

	// 新的序列化拷贝之后的对象实例
	return &newKeyWords

}

// keywords 关键字 map
type KeyWords map[string]*KeyWord

// 更新某些热点词

// 深拷贝
// 序列化拷贝所有的引用对象，得到一个新对象，空间地址不同
func (words KeyWords) DeepClone(updateWorks []*KeyWord) KeyWords {

	newWord := KeyWords{}

	for k, v := range words {

		newWord[k] = v.Clone()
	}

	return newWord

}

// 浅拷贝
// 只拷贝了地址，所有的引用对象都是同一个空间地址
func (words KeyWords) ShallowCopy(updateWorks []*KeyWord) KeyWords {

	newWord := KeyWords{}

	for k, v := range words {

		// 浅拷贝
		// 这里只拷贝了地址
		newWord[k] = v

	}

	return newWord

}
