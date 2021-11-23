package structural

import "testing"

func TestAdapter(t *testing.T){
	mysql := MakeMysql()
	mysql.Query()
	redis := MakeRedisToMysqlAdapter(&Redis{})
	redis.Query()
}
