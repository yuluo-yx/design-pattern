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
