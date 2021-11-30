package behavioral

import (
	"fmt"
	"reflect"
)

type PushSubscriber interface {
	Update(news string)
}

type PushPublisher interface {
	Attach(PushSubscriber)
	Detach()
	Notify()
}

type ConcretePushPublisher struct {
	subscribers []PushSubscriber
	lastNews    string
}

func (p *ConcretePushPublisher) Attach(s PushSubscriber) {
	p.subscribers = append(p.subscribers, s)
}

func (p *ConcretePushPublisher) Detach(s PushSubscriber) {
	for i, v := range p.subscribers{
		if reflect.DeepEqual(v, s){
			fmt.Printf("detach subscriber \n")
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
		}
	}
}

func (p *ConcretePushPublisher) Notify() {
	for _, v := range p.subscribers{
		v.Update(p.lastNews)
	}
}

func (p *ConcretePushPublisher) Publish(news string){
	p.lastNews = news
	p.Notify()
}

type ConcretePushSubscriber struct {
	name string
}

func (s *ConcretePushSubscriber) Update(news string) {
	fmt.Printf("%s get news '%s' from publisher\n", s.name, news)
}

func MakePushPublisher() *ConcretePushPublisher{
	return &ConcretePushPublisher{
		subscribers: make([]PushSubscriber, 0),
	}
}

func MakePushSubscriber(name string) *ConcretePushSubscriber{
	return &ConcretePushSubscriber{
		name: name,
	}
}
