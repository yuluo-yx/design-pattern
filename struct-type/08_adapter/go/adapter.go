package main

import (
	"encoding/json"
	"encoding/xml"
)

// 对象适配器，配置器持有遗留对象实例
// 如果实现类适配器，可以通过嵌入遗留对象实现，但不常用。
// 实现 json 数据接口并组合 xml 输出
type Adapter struct {
	legacy *LegacyDevice
}

func NewAdapter(legacy *LegacyDevice) *Adapter {

	return &Adapter{legacy: legacy}
}

func (a *Adapter) GetJSONData() string {

	xmlData := a.legacy.GetXMLData()

	var data Data
	xml.Unmarshal([]byte(xmlData), &data)
	jsonBytes, _ := json.Marshal(map[string]string{"value": data.Value})

	return string(jsonBytes)
}
