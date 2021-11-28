// 外观模式

// 优点
// 对使用方屏蔽子系统组件
// 子系统与客户之间的松耦合


// 缺点
// 增加新的子系统可能需要修改外观类，不符合开闭原则

package structural

import "fmt"

type CarComponent interface {
	Start()
}

type Engine struct {
}

func (*Engine) Start() {
	fmt.Println("Engine starting...")
}

type Turbo struct {
}

func (*Turbo) Start() {
	fmt.Println("Turbo starting...")
}

type GearBox struct {
}

func (*GearBox) Start() {
	fmt.Println("GearBox starting")
}

type Car struct {
	engine *Engine
	turbo *Turbo
	gearbox *GearBox
}

func (c *Car) Start()  {
	c.engine.Start()
	c.turbo.Start()
	c.gearbox.Start()
}
