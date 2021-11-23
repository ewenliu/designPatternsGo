package structural

import "testing"

func TestComposite(t *testing.T) {
	d1 := &Department{name: "sw"}
	d2 := &Department{name: "hw"}
	root := &Company{name: "nokia", sub: make([]BaseDepartment, 0)}
	root.AddSub(d1)
	root.AddSub(d2)

	d3 := &Department{name: "hz_sw"}
	d4 := &Department{name: "hz_hw"}
	hz := &Company{name: "nokia_hz", sub: make([]BaseDepartment, 0)}
	hz.AddSub(d3)
	hz.AddSub(d4)

	root.AddSub(hz)
	root.Display("")

	hz.Remove(d4)
	root.Display("")

	root.Remove(hz)
	root.Display("")

}
