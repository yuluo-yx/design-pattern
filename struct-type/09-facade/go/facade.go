package main

import (
	"fmt"
)

type ChooseProduct interface {
	// 选择产品
	Choose() any
}

type Payment interface {
	// 支付
	Pay() bool
}

type Repository interface {
	// 仓库打包
	Package() bool
}

type Transportation interface {
	// 物流运输
	Transit() bool
}

type ShoppingApp interface {
	// 购物 app 封装所有流程
	Shopping() bool
}

type ChooseProductImpl struct{}

func (c *ChooseProductImpl) Choose() any {
	fmt.Println("✓ 选择了商品：iPhone 15 Pro")
	return "iPhone 15 Pro"
}

type PaymentImpl struct{}

func (p *PaymentImpl) Pay() bool {
	fmt.Println("✓ 支付成功：￥8,999")
	return true
}

type RepositoryImpl struct{}

func (r *RepositoryImpl) Package() bool {
	fmt.Println("✓ 仓库完成打包")
	return true
}

type TransportationImpl struct{}

func (t *TransportationImpl) Transit() bool {
	fmt.Println("✓ 商品已发货，预计3天内送达")
	return true
}

// ===================== sub system =======================

// 购物APP门面
type ShoppingAppImpl struct {
	chooseProduct  ChooseProduct
	payment        Payment
	repository     Repository
	transportation Transportation
}

// 创建新的购物APP
func NewShoppingApp() ShoppingApp {

	// 依次管理依赖的 sub system 依赖关系
	return &ShoppingAppImpl{
		chooseProduct:  &ChooseProductImpl{},
		payment:        &PaymentImpl{},
		repository:     &RepositoryImpl{},
		transportation: &TransportationImpl{},
	}
}

// Shopping 封装了整个购物流程
func (s *ShoppingAppImpl) Shopping() bool {
	fmt.Println("=== 开始购物 ===")

	// 1. 选择商品
	product := s.chooseProduct.Choose()
	if product == nil {
		fmt.Println("❌ 商品选择失败")
		return false
	}

	// 2. 支付
	if !s.payment.Pay() {
		fmt.Println("❌ 支付失败")
		return false
	}

	// 3. 仓库打包
	if !s.repository.Package() {
		fmt.Println("❌ 仓库打包失败")
		return false
	}

	// 4. 物流配送
	if !s.transportation.Transit() {
		fmt.Println("❌ 物流配送失败")
		return false
	}

	fmt.Println("=== 购物成功完成 ===")

	return true
}
