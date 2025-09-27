package main

import "fmt"

// ProjectVisitor 项目访问者接口
// 定义了对不同员工类型的访问操作
type ProjectVisitor interface {
	VisitTester(tester *Tester)
	VisitDeveloper(developer *Developer)
	VisitHRBP(hrbp *HRBP)
}

// Employee 员工接口 - 元素接口
// 定义了接受项目访问者的方法
type Employee interface {
	Accept(visitor ProjectVisitor)
	GetName() string
	GetPosition() string
}

// Tester 测试员工结构体 - 具体元素
type Tester struct {
	Name string
}

// NewTester 创建测试员工
func NewTester(name string) *Tester {

	return &Tester{Name: name}
}

// Accept 实现Employee接口 - 访问者模式的关键
func (t *Tester) Accept(visitor ProjectVisitor) {

	// 这里实现了双分派的第一部分
	visitor.VisitTester(t)
}

func (t *Tester) GetName() string {
	return t.Name
}

func (t *Tester) GetPosition() string {
	return "测试工程师"
}

// ExecuteTest 测试员工特有的方法 - 执行测试
func (t *Tester) ExecuteTest() {
	fmt.Printf("%s 正在执行测试用例...\n", t.Name)
}

// SubmitBugReport 测试员工特有的方法 - 提交Bug报告
func (t *Tester) SubmitBugReport() {
	fmt.Printf("%s 提交了Bug报告\n", t.Name)
}

// Developer 开发员工结构体 - 具体元素
type Developer struct {
	Name string
}

// NewDeveloper 创建开发员工
func NewDeveloper(name string) *Developer {
	return &Developer{Name: name}
}

// Accept 实现Employee接口
func (d *Developer) Accept(visitor ProjectVisitor) {
	visitor.VisitDeveloper(d)
}

func (d *Developer) GetName() string {
	return d.Name
}

func (d *Developer) GetPosition() string {
	return "开发工程师"
}

// WriteCode 开发员工特有的方法 - 编写代码
func (d *Developer) WriteCode() {
	fmt.Printf("%s 正在疯狂敲代码...\n", d.Name)
}

// FixBug 开发员工特有的方法 - 修复Bug
func (d *Developer) FixBug() {
	fmt.Printf("%s 修复了一个Bug（可能引入了两个新Bug）\n", d.Name)
}

// HRBP HR员工结构体 - 具体元素
type HRBP struct {
	Name string
}

// NewHRBP 创建HR员工
func NewHRBP(name string) *HRBP {
	return &HRBP{Name: name}
}

// Accept 实现Employee接口
func (h *HRBP) Accept(visitor ProjectVisitor) {
	visitor.VisitHRBP(h)
}

func (h *HRBP) GetName() string {
	return h.Name
}

func (h *HRBP) GetPosition() string {
	return "人力资源业务伙伴"
}

// Recruit HR员工特有的方法 - 招聘
func (h *HRBP) Recruit() {
	fmt.Printf("%s 正在招聘新员工（画饼中...）\n", h.Name)
}

// OrganizeActivity HR员工特有的方法 - 组织活动
func (h *HRBP) OrganizeActivity() {
	fmt.Printf("%s 组织了团建活动\n", h.Name)
}
