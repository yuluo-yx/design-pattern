package main

import (
	"fmt"
)

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
