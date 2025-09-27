package main

import "fmt"

// IMProjectVisitor IM项目访问者 - 具体访问者
// 实现IM项目中各种员工的工作分配
type IMProjectVisitor struct{}

// NewIMProjectVisitor 创建IM项目访问者
func NewIMProjectVisitor() *IMProjectVisitor {
	return &IMProjectVisitor{}
}

func (v *IMProjectVisitor) VisitTester(tester *Tester) {

	fmt.Printf("=== IM项目中的%s ===\n", tester.GetName())
	fmt.Printf("%s 在IM项目中的工作：当陪聊机器人\n", tester.GetName())

	tester.ExecuteTest()
	fmt.Println()
}

func (v *IMProjectVisitor) VisitDeveloper(developer *Developer) {

	fmt.Printf("=== IM项目中的%s ===\n", developer.GetName())
	fmt.Printf("%s 在IM项目中的工作：开发微信App\n", developer.GetName())

	developer.WriteCode()
	fmt.Println()
}

func (v *IMProjectVisitor) VisitHRBP(hrbp *HRBP) {

	fmt.Printf("=== IM项目中的%s ===\n", hrbp.GetName())
	fmt.Printf("%s 在IM项目中的工作：宣传运营\n", hrbp.GetName())

	hrbp.OrganizeActivity()
	fmt.Println()
}

// ECommerceProjectVisitor 电商项目访问者 - 具体访问者
// 实现电商项目中各种员工的工作分配
type ECommerceProjectVisitor struct{}

// NewECommerceProjectVisitor 创建电商项目访问者
func NewECommerceProjectVisitor() *ECommerceProjectVisitor {
	return &ECommerceProjectVisitor{}
}

func (v *ECommerceProjectVisitor) VisitTester(tester *Tester) {

	fmt.Printf("=== 电商项目中的%s ===\n", tester.GetName())
	fmt.Printf("%s 在电商项目中的工作：客服小姐姐\n", tester.GetName())

	tester.SubmitBugReport()
	fmt.Println()
}

func (v *ECommerceProjectVisitor) VisitDeveloper(developer *Developer) {

	fmt.Printf("=== 电商项目中的%s ===\n", developer.GetName())
	fmt.Printf("%s 在电商项目中的工作：开发淘宝App\n", developer.GetName())

	developer.WriteCode()
	developer.FixBug()
	fmt.Println()
}

func (v *ECommerceProjectVisitor) VisitHRBP(hrbp *HRBP) {

	fmt.Printf("=== 电商项目中的%s ===\n", hrbp.GetName())
	fmt.Printf("%s 在电商项目中的工作：直播带货主播\n", hrbp.GetName())

	hrbp.Recruit()
	fmt.Println()
}

// GameProjectVisitor 游戏项目访问者 - 演示扩展性
type GameProjectVisitor struct{}

// NewGameProjectVisitor 创建游戏项目访问者
func NewGameProjectVisitor() *GameProjectVisitor {
	return &GameProjectVisitor{}
}

func (v *GameProjectVisitor) VisitTester(tester *Tester) {

	fmt.Printf("=== 游戏项目中的%s ===\n", tester.GetName())
	fmt.Printf("%s 在游戏项目中的工作：游戏陪玩\n", tester.GetName())

	tester.ExecuteTest()
	fmt.Println()
}

func (v *GameProjectVisitor) VisitDeveloper(developer *Developer) {

	fmt.Printf("=== 游戏项目中的%s ===\n", developer.GetName())
	fmt.Printf("%s 在游戏项目中的工作：开发王者荣耀\n", developer.GetName())

	developer.WriteCode()
	fmt.Println()
}

func (v *GameProjectVisitor) VisitHRBP(hrbp *HRBP) {

	fmt.Printf("=== 游戏项目中的%s ===\n", hrbp.GetName())
	fmt.Printf("%s 在游戏项目中的工作：游戏推广大使\n", hrbp.GetName())

	hrbp.OrganizeActivity()
	fmt.Println()
}
