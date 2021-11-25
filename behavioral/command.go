package behavioral

import "fmt"

type Command interface {
	Execute()
}

type Light struct {

}

func (*Light) TurnOn()  {
	fmt.Println("Light turn on")
}

func (*Light) TurnOff()  {
	fmt.Println("Light turn off")
}

type FlipUp struct {
	light *Light
}

type FlipDown struct {
	light *Light
}

func (f *FlipUp ) Execute()  {
	f.light.TurnOn()
}

func (f *FlipDown ) Execute()  {
	f.light.TurnOff()
}

type LightSwitch struct {
	fu *FlipUp
	fd *FlipDown
}

func (ls *LightSwitch ) Switch(cmd string)  {
	switch cmd {
	case "ON":
		ls.fu.Execute()
	case "OFF":
		ls.fd.Execute()
	default:
		fmt.Println("No operation")
	}
}

func MakeLightSwitch() *LightSwitch{
	light := &Light{}
	fu := &FlipUp{light: light}
	fd := &FlipDown{light: light}
	return &LightSwitch{fd: fd, fu: fu}
}
