package main

import "fmt"

// BlackHeartCompany 黑心小公司结构体 - 对象结构
// 管理所有员工，并支持访问者模式遍历
type BlackHeartCompany struct {
	employees   []Employee
	companyName string
}

// NewBlackHeartCompany 创建黑心小公司
func NewBlackHeartCompany(companyName string) *BlackHeartCompany {

	return &BlackHeartCompany{
		companyName: companyName,
		employees:   make([]Employee, 0),
	}
}

// AddEmployee 添加员工（招聘）
func (c *BlackHeartCompany) AddEmployee(employee Employee) {

	c.employees = append(c.employees, employee)

	fmt.Printf("恭喜 %s（%s）加入%s大家庭！\n",
		employee.GetName(), employee.GetPosition(), c.companyName)
}

// RemoveEmployee 移除员工（裁员）
func (c *BlackHeartCompany) RemoveEmployee(employee Employee) {

	for i, emp := range c.employees {
		if emp == employee {
			c.employees = append(c.employees[:i], c.employees[i+1:]...)
			fmt.Printf("%s 被优化了，祝工作顺利！\n", employee.GetName())
			break
		}
	}
}

// StartProject 启动项目 - 访问者模式的核心应用
// 让所有员工接受项目访问者，分配不同的工作
func (c *BlackHeartCompany) StartProject(projectVisitor ProjectVisitor, projectName string) {
	
	fmt.Printf("🚀 %s 启动 %s 项目！\n", c.companyName, projectName)
	fmt.Println("老板：兄弟们，新项目来了，大家各司其职！")
	fmt.Println("================================================")

	// 遍历所有员工，让他们接受项目访问者
	for _, employee := range c.employees {
		employee.Accept(projectVisitor)
	}

	fmt.Println("================================================")
	fmt.Println("老板：很好，大家都有活干了！记住，deadline是一个月！")
	fmt.Println()
}

// GetEmployeeCount 获取员工总数
func (c *BlackHeartCompany) GetEmployeeCount() int {

	return len(c.employees)
}

// ShowCompanyInfo 显示公司概况
func (c *BlackHeartCompany) ShowCompanyInfo() {

	fmt.Printf("=== %s 公司概况 ===\n", c.companyName)
	fmt.Printf("员工总数：%d 人\n", len(c.employees))
	fmt.Println("员工列表：")
	
	for _, employee := range c.employees {
		fmt.Printf("- %s（%s）\n", employee.GetName(), employee.GetPosition())
	}

	fmt.Println()
}
