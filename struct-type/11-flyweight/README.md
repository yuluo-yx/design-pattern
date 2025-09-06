# 享元模式 Flyweight Design Pattern）

享元模式应用最广的就是池化技术。例如 String 常量池，数据库连接池等。

从池化技术不难猜到，享元模式的用途主要是减少对象的创建次数，以减少内存占用和提高性能。

### 享元模式状态

先来看定义：

内部状态：对象的常量数据被称为内在状态，位于对象内部，只能访问不能修改；
外部状态：对象的其他状态常常被其他对象“从外部”改变，而发生变化。

举个例子：

我今天出门了，天气下雨了，我被淋湿了，我去服装店买了身新衣服换上。

这个例子中 “我” 这个对象状态：
内部状态：姓名，性别，生日 等；不可变只可读；
外部状态：雨水淋湿，衣服变换，分别对天气对象和服装店对象等外部对象影响而发生状态变化。

接下来思考一个问题，🤔 这些对象状态被保存在哪里？

继续以上面的例子为例：

大范围中：处于地球上；
小范围中：处于 xxx 市 xxx 区。

那便可以以此来作为享元模式的容器对象存储对象的外部状态。

### 使用场景

各种池化场景：线程池，连接池等。

## 1. 应用

当系统中存在大量对象且这些对象是重量级对象。

### 1.1 优点

减少对象的创建次数，复用对象。

### 1.2 缺点

1. 代码变得更加复杂，拆分了实体状态；
2. 牺牲了执行时间，需要重新渲染外部状态。

## 2. 实现分析

以 Golang 作为代码实现分析：

```go
package main

// 享元对象：Person（只存内部状态）
type Person struct {
	name   string // 内部状态
	gender string // 内部状态
	birth  string // 内部状态
}

// 外部状态
type ExternalState struct {
	weather   string
	clothes   string
	areaBig   string
	areaSmall string
}

// 外部状态容器（全局管理外部状态）
// 外部状态容器相比对象的内在状态更容易变化，因此可以单独存储和管理
// 且占用内存更少，适合频繁变化的状态
type ExternalStateManager struct {
	stateMap map[string]*ExternalState
}

func NewExternalStateManager() *ExternalStateManager {
	return &ExternalStateManager{stateMap: make(map[string]*ExternalState)}
}

func (m *ExternalStateManager) SetState(personKey string, state *ExternalState) {
	m.stateMap[personKey] = state
}

func (m *ExternalStateManager) GetState(personKey string) *ExternalState {
	return m.stateMap[personKey]
}

// 享元工厂
type PersonFactory struct {
	persons map[string]*Person
}

func NewPersonFactory() *PersonFactory {

	// 注意：享元对象的状态只能在创建时引入，不对外提供修改接口
	// Java 中是构造方法传入
	return &PersonFactory{persons: make(map[string]*Person)}
}

func (f *PersonFactory) GetPerson(name, gender, birth string) *Person {

	key := name + gender + birth
	if p, ok := f.persons[key]; ok {
		return p
	}
	p := &Person{name, gender, birth}
	f.persons[key] = p

	return p
}
```

客户端实现：

```go
func main() {
	// 享元工厂
	personFactory := NewPersonFactory()
	// 外部状态容器
	stateManager := NewExternalStateManager()

	// 创建“我”对象
	me := personFactory.GetPerson("小明", "男", "1990-01-01")
	personKey := "小明男1990-01-01"

	// 场景1：下雨被淋湿
	stateManager.SetState(personKey, &ExternalState{
		weather:   "下雨",
		clothes:   "新衣服",
		areaBig:   "地球",
		areaSmall: "上海市浦东新区",
	})
	state := stateManager.GetState(personKey)
	fmt.Printf("%s(%s, %s) 今天出门了，处于%s/%s，天气%s，我被淋湿了，我去服装店买了身%s换上。 [me 对象: %p]\n",
		me.name, me.gender, me.birth, state.areaBig, state.areaSmall, state.weather, state.clothes, me)

	// 场景2：打球流汗
	stateManager.SetState(personKey, &ExternalState{
		weather:   "晴天",
		clothes:   "运动服",
		areaBig:   "地球",
		areaSmall: "上海市徐汇区",
	})
	state2 := stateManager.GetState(personKey)
	fmt.Printf("%s(%s, %s) 去打球，处于%s/%s，天气%s，打球流汗了，我去更衣室换了身%s。 [me 对象: %p]\n",
		me.name, me.gender, me.birth, state2.areaBig, state2.areaSmall, state2.weather, state2.clothes, me)
}
```

输出如下：

```markdown
小明(男, 1990-01-01) 今天出门了，处于地球/上海市浦东新区，天气下雨，我被淋湿了，我去服装店买了身新衣服换上。 [me 对象: 0x1400007e090]
小明(男, 1990-01-01) 去打球，处于地球/上海市徐汇区，天气晴天，打球流汗了，我去更衣室换了身运动服。 [me 对象: 0x1400007e090]
```
