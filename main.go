package main

// tạo đối tượng có nhiều thuộc tính
import (
	"DPbuilder/builder"
	"fmt"
)

func main() {

	normalBuilder := builder.GetBuilder("normal")
	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Printf("normal house doortype: %s\n", normalHouse.GetDoorType())

	iglooBuilder := builder.GetBuilder("igloo")
	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Printf("Igloo house door type: %s\n", iglooHouse.GetDoorType())

}
