package structural

import (
	"fmt"
	"reflect"
)

type BaseDepartment interface {
	Display(company string)
}

type Department struct {
	name string
}

func (d *Department ) Display(company string)  {
	fmt.Printf("this is %s %s \n", company, d.name)
}

type Company struct {
	sub []BaseDepartment
	name string
}

func (c *Company) AddSub(o BaseDepartment)  {
	c.sub = append(c.sub, o)
}

func (c *Company ) Remove(o BaseDepartment)  {
	removeIndex := -1
	for k, v := range c.sub{
		if reflect.DeepEqual(v, o){
			removeIndex = k
		}
	}
	if removeIndex != -1{
		c.sub = append(c.sub[:removeIndex], c.sub[removeIndex+1:]...)
	}
}

func (c *Company ) Display(company string)  {
	for _, s := range c.sub{
		s.Display(c.name)
	}
}