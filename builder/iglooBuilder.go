package builder

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}
func (iglooBuilder *iglooBuilder) setWindowType() {
	iglooBuilder.windowType = "snow window"
}
func (iglooBuilder *iglooBuilder) setDoorType() {
	iglooBuilder.doorType = "snow door"
}
func (iglooBuilder *iglooBuilder) setNumFloor() {
	iglooBuilder.floor = 2
}
func (iglooBuilder *iglooBuilder) getHouse() House {
	return House{
		doorType:   iglooBuilder.doorType,
		windowType: iglooBuilder.windowType,
		floor:      iglooBuilder.floor,
	}
}
