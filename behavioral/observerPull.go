package behavioral

import (
	"fmt"
	"reflect"
)

type PullSubscriber interface {
	Update()
}

type PullPublisher interface {
	Attach(PullSubscriber)
	Detach()
	Notify()
}

type ConcretePullPublisher struct {
	subscribers []PullSubscriber
	lastNews    string
}

func (p *ConcretePullPublisher) Attach(s PullSubscriber) {
	p.subscribers = append(p.subscribers, s)
}

func (p *ConcretePullPublisher) Detach(s PullSubscriber) {
	for i, v := range p.subscribers{
		if reflect.DeepEqual(v, s){
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
		}
	}
}

func (p *ConcretePullPublisher) Notify() {
	for _, v := range p.subscribers{
		v.Update()
	}
}

func (p *ConcretePullPublisher) Publish(news string){
	p.lastNews = news
	p.Notify()
}

type ConcretePullSubscriber struct {
	publisher *ConcretePullPublisher
	name string
}

func (s *ConcretePullSubscriber) Update() {
	fmt.Printf("%s get news '%s' from publisher\n", s.name, s.publisher.lastNews)
}

func MakePullPublisher() *ConcretePullPublisher{
	return &ConcretePullPublisher{
		subscribers: make([]PullSubscriber, 0),
	}
}

func MakePullSubscriber(name string) *ConcretePullSubscriber{
	return &ConcretePullSubscriber{
		name: name,
	}
}
