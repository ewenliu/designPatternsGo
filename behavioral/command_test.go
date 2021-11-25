package behavioral

import "testing"

func TestCommand(t *testing.T){
	ls := MakeLightSwitch()
	ls.Switch("ON")
	ls.Switch("OFF")
}
