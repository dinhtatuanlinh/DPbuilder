package builder

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}
func (normalBuilder *normalBuilder) setWindowType() {
	normalBuilder.windowType = "wooden window"
}
func (normalBuilder *normalBuilder) setDoorType() {
	normalBuilder.doorType = "wooden door"
}
func (normalBuilder *normalBuilder) setNumFloor() {
	normalBuilder.floor = 2
}
func (normalBuilder *normalBuilder) getHouse() House {
	return House{
		doorType:   normalBuilder.doorType,
		windowType: normalBuilder.windowType,
		floor:      normalBuilder.floor,
	}
}
