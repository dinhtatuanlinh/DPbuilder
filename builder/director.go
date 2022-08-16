package builder

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}
func (director *Director) SetBuilder(b IBuilder) {
	director.builder = b
}
func (director *Director) BuildHouse() House {
	director.builder.setDoorType()
	director.builder.setNumFloor()
	director.builder.setWindowType()
	return director.builder.getHouse()
}
