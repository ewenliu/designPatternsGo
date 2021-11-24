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