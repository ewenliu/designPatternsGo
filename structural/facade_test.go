package structural

import "testing"

func TestFacade(t *testing.T) {
	engine := &Engine{}
	turbo := &Turbo{}
	gearbox := &GearBox{}
	car := &Car{engine: engine, turbo: turbo, gearbox: gearbox}
	car.Start()
}
