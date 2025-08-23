package main

import "fmt"

func main() {

	legacy := &LegacyDevice{}
	adapter := NewAdapter(legacy)

	fmt.Println("Legacy XML:", legacy.GetXMLData())
	fmt.Println("============ adapter output ===========")
	fmt.Println("Adapted JSON:", adapter.GetJSONData())
}
