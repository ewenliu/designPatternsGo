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
