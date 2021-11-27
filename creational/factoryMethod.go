// 优点
// 系统中加入新产品时，无须修改抽象工厂和抽象产品提供的接口，符合开闭原则

// 缺点
// 添加新产品时，除了需要编写新的具体产品类，而且还要编写与之对应的具体工厂类


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
