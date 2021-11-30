package behavioral

import "testing"

func TestPushObserver(t *testing.T)  {
	publisher := MakePushPublisher()
	sub1 := MakePushSubscriber("s1")
	sub2 := MakePushSubscriber("s2")
	publisher.Attach(sub1)
	publisher.Attach(sub2)
	publisher.Publish("hello everyone!")
	publisher.Detach(sub1)
	publisher.Publish("who are still alive!")
}

func TestPullObserver(t *testing.T)  {
	publisher := MakePullPublisher()
	sub1 := MakePullSubscriber("s1")
	sub2 := MakePullSubscriber("s2")
	sub1.publisher = publisher
	sub2.publisher = publisher
	publisher.Attach(sub1)
	publisher.Attach(sub2)
	publisher.Publish("hello everyone!")
	publisher.Detach(sub1)
	publisher.Publish("who are still alive!")
}
