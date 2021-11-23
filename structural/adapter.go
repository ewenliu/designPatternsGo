package structural

import "fmt"

type Mysql struct {

}

func (*Mysql) Query()  {
	fmt.Println("Mysql query sth")
}

func MakeMysql() *Mysql{
	return &Mysql{}
}

type Redis struct {

}

func (*Redis) Execute() {
	fmt.Println("Redis do executing")
}

type RedisToMysqlAdapter struct {
	redis *Redis
}

func (ra *RedisToMysqlAdapter ) Query()  {
	ra.redis.Execute()
}

func MakeRedisToMysqlAdapter(redis *Redis) *RedisToMysqlAdapter  {
	return &RedisToMysqlAdapter{
		redis: redis,
	}
}
