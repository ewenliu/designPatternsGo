package creational

import "testing"

func TestFactoryMethod(t *testing.T){
	var af AnimalFactory

	af = &DogFactory{}
	dog := af.GetAnimal()
	dog.SetInfo("lucky")
	dog.Speak()

	af = &CatFactory{}
	cat := af.GetAnimal()
	cat.SetInfo("lucy")
	cat.Speak()

	af = &BirdFactory{}
	bird := af.GetAnimal()
	bird.SetInfo("sin")
	bird.Speak()
}
