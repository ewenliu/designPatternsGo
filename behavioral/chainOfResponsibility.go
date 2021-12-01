// 责任链模式

// 优点
// 1、将请求者与处理者分离，请求者并不知道请求是被哪个处理者所处理，易于扩展
// 应用场景：若一个请求可能由一个对请求有链式优先级的处理群所处理时，可以考虑责任链模式
// 如银行的客户请求处理系统也可以用责任链模式实现

// 缺点
// 责任链比较长，会有比较大的性能问题
// 责任链比较长，若业务出现问题，比较难定位是哪个处理者的问题


package behavioral

import "fmt"

type Manger interface {
	setSuperiorManager(superior Manger)
	handleRequest(day int)
}

type LineManager struct {
	superior Manger
}

func (lm *LineManager ) setSuperiorManager(superior Manger)  {
	lm.superior = superior
}

func (lm *LineManager ) handleRequest(day int)  {
	if day <= 3{
		fmt.Println("Line manager accept, approved!")
		return
	}
	fmt.Println("Line manager can't handle, transfer to superior manger!")
	lm.superior.handleRequest(day)
}

type SeniorManger struct {
	superior Manger
}

func (sm *SeniorManger ) setSuperiorManager(superior Manger)  {
	sm.superior = superior
}

func (sm *SeniorManger ) handleRequest(day int)  {
	if day >3 && day <= 7{
		fmt.Println("Senior manager accept, approved!")
		return
	}
	fmt.Println("Senior manager can't handle, transfer to superior manger!")
	sm.superior.handleRequest(day)
}

type Chief struct {
	superior Manger
}

func (cf *Chief ) setSuperiorManager(superior Manger)  {
	cf.superior = superior
}

func (cf *Chief ) handleRequest(day int)  {
	fmt.Println("Chief accept, approved!")
}
