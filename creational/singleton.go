package creational

import "sync"

type DBPool struct {
}

var db *DBPool

var once sync.Once

func GetDB() *DBPool {
	once.Do(func() {
		db = &DBPool{}
	})
	return db
}
