package main

import (
	"fmt"
	"time"
)

func main() {

	updateAt, _ := time.Parse("2006", "2020")

	oldWords := KeyWords{
		"golang": &KeyWord{
			Word:     "golang",
			Visit:    100,
			UpdateAt: &updateAt,
		},
		"python": &KeyWord{
			Word:     "python",
			Visit:    200,
			UpdateAt: &updateAt,
		},
	}
	now := time.Now()
	updatedWords := []*KeyWord{
		{
			Word:     "python",
			Visit:    10,
			UpdateAt: &now,
		},
	}

	fmt.Println("浅拷贝示例，经过拷贝的 time 指针地址相同：============================================")
	newWords1 := oldWords.ShallowCopy(updatedWords)
	fmt.Printf("原始对象实例和拷贝之后对象实例：oldWords: %p, newWords: %p\n", &oldWords, &newWords1)
	fmt.Printf("原始对象实例和拷贝之后对象实例的引用字段：oldWords['python'].time: %p, newWords['python'].time: %p\n", &oldWords["python"].UpdateAt, &newWords1["python"].UpdateAt)

	fmt.Println("\n深拷贝示例，经过拷贝的 time 指针地址不同：============================================")
	newWords2 := oldWords.DeepClone(updatedWords)
	fmt.Printf("原始对象实例和拷贝之后对象实例：oldWords: %p, newWords: %p\n", &oldWords, &newWords2)
	fmt.Printf("原始对象实例和拷贝之后对象实例的引用字段：oldWords['python'].time: %p, newWords['python'].time: %p\n", &oldWords["python"].UpdateAt, &newWords2["python"].UpdateAt)

}
