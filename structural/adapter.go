// 适配器模式

// 优点
// 目标类和适配者类解耦，通过引入一个适配器类来重用现有的适配者类，而无须修改原有代码，符合开闭原则

// 缺点
// 当被适配的类方法名更改时，适配器也要同步更改


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
