package creational

import "fmt"

type Speaker interface {
	Speak()
	SetInfo(name string)
}

type Animal struct {
	name string
}

func (a *Animal) SetInfo(name string) {
	a.name = name
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

type Bird struct {
	Animal
}

func (d *Dog) Speak() {
	fmt.Println("a wu a wu")
}

func (c *Cat) Speak() {
	fmt.Println("miao miao miao")
}

func (b *Bird) Speak() {
	fmt.Println("zhi zhi zhi")
}

type AnimalFactory interface {
	GetAnimal() Speaker
}

type DogFactory struct {
}

type CatFactory struct {
}
type BirdFactory struct {
}

func (df *DogFactory) GetAnimal() Speaker {
	return &Dog{}
}

func (cf *CatFactory) GetAnimal() Speaker {
	return &Cat{}
}

func (bf *BirdFactory) GetAnimal() Speaker {
	return &Bird{}
}
