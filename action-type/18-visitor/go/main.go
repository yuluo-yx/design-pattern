package main

import "fmt"

/*
访问者模式演示 - 黑心小公司的项目分工

访问者模式的核心思想：
1. 将数据结构（员工）与操作（项目工作）分离
2. 通过接口分派机制实现不同类型员工在不同项目中的不同行为
3. 符合开闭原则：新增项目类型无需修改员工结构体

注意：Go语言没有方法重载，所以用不同的方法名来实现"分派"效果（这里只是演示，Go 不需要此结构，可以使用反射或者 generate 来完成）
*/

func main() {

	// 创建黑心小公司
	company := NewBlackHeartCompany("996创新科技有限公司")

	// 招聘员工（公司的数据结构）
	company.AddEmployee(NewTester("小王"))
	company.AddEmployee(NewDeveloper("小李"))
	company.AddEmployee(NewHRBP("小张"))

	fmt.Println()
	company.ShowCompanyInfo()

	// 创建不同的项目访问者（不同的算法）
	imProject := NewIMProjectVisitor()
	ecommerceProject := NewECommerceProjectVisitor()
	gameProject := NewGameProjectVisitor()

	// 演示访问者模式：同样的员工，不同的项目，不同的工作安排
	fmt.Println("📱 三个月前，老板跟风做IM:")
	company.StartProject(imProject, "仿微信IM")

	fmt.Println("🛒 现在，老板又想做电商:")
	company.StartProject(ecommerceProject, "仿淘宝电商")

	// 演示访问者模式的扩展性
	fmt.Println("🎮 老板脑洞大开，又想做游戏了:")
	company.StartProject(gameProject, "仿王者荣耀手游")

	// Go 版本的函数式特性
	fmt.Println("\n=== Go 函数式访问者 ===")
	demonstrateFunctionalVisitor(company)
}

// demonstrateFunctionalVisitor 演示 Go 版本特有的函数式访问者
// 由于没有函数重载，只能通过断言来根据不同的类型调用不同函数
// 没有访问者模式，也可以这么操作。
// 随着角色越来越多，会越复杂和庞大。也可以用反射或者 code generate
func demonstrateFunctionalVisitor(company *BlackHeartCompany) {
	// 使用函数作为访问者
	blockchainVisitor := func(employee Employee) {
		fmt.Printf("🔗 区块链项目中的%s: ", employee.GetName())
		switch emp := employee.(type) {
		case *Tester:
			fmt.Println("挖矿测试员 - 测试挖矿算法效率")
		case *Developer:
			fmt.Println("智能合约开发者 - 开发去中心化应用")
		case *HRBP:
			fmt.Println("币圈KOL - 推广数字货币")
		default:
			fmt.Printf("未知员工类型: %T\n", emp)
		}
	}

	fmt.Println("使用函数式访问者访问所有员工:")
	visitAllEmployees(company, blockchainVisitor)
}

// visitAllEmployees 使用函数作为访问者遍历所有员工
func visitAllEmployees(company *BlackHeartCompany, visitor func(Employee)) {
	// 这里展示了Go的反射和类型断言特性
	// 虽然不是标准的访问者模式，但展示了Go的灵活性
	for i := 0; i < company.GetEmployeeCount(); i++ {
		if i < len(company.employees) {
			visitor(company.employees[i])
		}
	}
	fmt.Println()
}
