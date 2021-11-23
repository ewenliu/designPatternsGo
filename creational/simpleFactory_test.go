package creational

import "testing"

func TestSimpleFactory(t *testing.T) {
	car := ChooseBicycle("car")
	car.Work()

	bicycle := ChooseBicycle("bicycle")
	bicycle.Work()

	airplane := ChooseBicycle("airplane")
	airplane.Work()
}
