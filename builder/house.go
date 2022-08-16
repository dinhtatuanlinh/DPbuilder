package builder

// house struct
type House struct {
	windowType string
	doorType   string
	floor      int
}

// get windowType
func (house *House) GetWindowType() string {
	return house.windowType
}

func (house *House) GetDoorType() string {
	return house.doorType
}
func (house *House) GetnumFloor() int {
	return house.floor
}
