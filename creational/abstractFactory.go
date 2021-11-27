package creational

import "fmt"
// 优点
// 加入新产品工厂时很方便，符合开闭原则

// 缺点
// 添加新产品时，基本每个产品工厂都要涉及到改变


type Factory interface {
	ProduceTV() TV
	ProduceMobilePhone() MobilePhone
}

type TV interface {
	Show()
}

type MobilePhone interface {
	Call()
}

type AppleTV struct {
}

func (tv *AppleTV) Show() {
	fmt.Println("Apple TV showing.....")
}

type HuaweiTV struct {
}

func (tv *HuaweiTV) Show() {
	fmt.Println("Huawei TV showing.....")
}

type IPhone struct {
}

func (phone *IPhone) Call() {
	fmt.Println("IPhone calling")
}

type MateX struct {
}

func (phone *MateX) Call() {
	fmt.Println("Matex calling")
}

type HuaweiFactory struct {
}

func (hf *HuaweiFactory) ProduceTV() TV {
	return &HuaweiTV{}
}

func (hf *HuaweiFactory) ProduceMobilePhone() MobilePhone {
	return &MateX{}
}

type AppleFactory struct {
}

func (apple *AppleFactory) ProduceTV() TV {
	return &AppleTV{}
}

func (apple *AppleFactory) ProduceMobilePhone() MobilePhone {
	return &IPhone{}
}
