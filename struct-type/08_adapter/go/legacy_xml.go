package main

import "encoding/xml"

type LegacyDevice struct{}

type Data struct {
	XMLName xml.Name `xml:"data"`
	Value   string   `xml:"value"`
}

func (d *LegacyDevice) GetXMLData() string {

	data := Data{Value: "Hello from legacy device"}
	bytes, _ := xml.Marshal(data)

	return string(bytes)
}
