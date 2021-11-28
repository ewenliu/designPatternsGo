// 组合模式

// 优点
// 清楚地定义分层次的复杂对象，表示对象的全部或部分层次，使得增加新构件也更容易

// 缺点
// 设计变得过于抽象，代码难以理解，如果业务规则很复杂，实现比较难以实现

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