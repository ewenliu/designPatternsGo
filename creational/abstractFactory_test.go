package creational

import "testing"

func TestAbstractFactory(t *testing.T) {
	apple := &AppleFactory{}
	iphone := apple.ProduceMobilePhone()
	iphone.Call()
	appleTV := apple.ProduceTV()
	appleTV.Show()

	huawei := &HuaweiFactory{}
	matex := huawei.ProduceMobilePhone()
	matex.Call()
	huaweiTV := huawei.ProduceTV()
	huaweiTV.Show()
}
