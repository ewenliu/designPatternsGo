// 优点
// 可以在调用时将所传入的参数保存在XML等格式的配置文件中，修改参数时无须修改任何源代码

// 缺点
// 工厂类的职责相对过重，增加新的产品需要修改工厂类的判断逻辑，这一点与开闭原则是相违背的


package creational

import (
	"fmt"
	"strings"
)

type Vehicle interface {
	Work()
}

type Car struct {
}

func (c *Car) Work() {
	fmt.Println("Car has four wheels")
}

type Bicycle struct {
}

func (b *Bicycle) Work() {
	fmt.Println("Bicycle has two wheels")
}

type AirPlane struct {
}

func (a *AirPlane) Work() {
	fmt.Println("Airplane doesn't has wheels but it can fly")
}

func ChooseBicycle(vehicle string) Vehicle {
	switch strings.ToUpper(vehicle){
	case "CAR":
		return new(Car)
	case "BICYCLE":
		return new(Bicycle)
	case "AIRPLANE":
		return new(AirPlane)
	default:
		panic("No such vehicle!")
	}
}
