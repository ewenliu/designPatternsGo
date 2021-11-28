// 命令模式

// 优点
// 分离了集合对象的遍历行为

// 缺点
// 类的个数成对增加，即一个container需要一个迭代器，如果增加多一种container，则需要多写一个迭代器，增加了系统的复杂性


package behavioral

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Container interface {
	GetIterator() Iterator
}

type NameIterator struct {
	names []string
	index int
}

func (ni *NameIterator) HasNext() bool {
	if ni.index < len(ni.names) {
		return true
	}
	return false
}

func (ni *NameIterator) Next() interface{} {
	if ni.HasNext() {
		res := ni.names[ni.index]
		ni.index++
		return res
	}
	return ni
}

type NameRepo struct {

}

func (*NameRepo) GetIterator() Iterator {
	nameIterator := new(NameIterator)
	nameIterator.names = []string{"xiaohei", "xiaobai"}
	return nameIterator
}

