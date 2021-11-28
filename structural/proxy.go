// 代理模式

// 优点
// 以控制对真实对象的使用权限
// 子系统与客户之间的松耦合


// 缺点
// 由于增加了代理对象，可能会造成请求的处理速度变慢


package structural

import "fmt"

type Handler interface {
	Get()
	Set()
}

type DB struct {
}

func (*DB) Get() {
	fmt.Println("this is real handler Get!")
}

func (*DB) Set() {
	fmt.Println("this is real handler Set!")
}

type DBProxy struct {
	Auth bool
	DB *DB
}

func (db *DBProxy ) Get()  {
	db.DB.Get()
}

func (db *DBProxy ) Set()  {
	if !db.Auth {
		fmt.Println("No permission")
		return
	}
	db.DB.Set()
}