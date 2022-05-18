package config

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RdbContext = context.Background()
var Rdb *redis.Client

func init() {
	port := AppConfig.GetString("redis.port")
	password := AppConfig.GetString("redis.passowrd")
	db := AppConfig.GetInt("DB")
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:" + port,
		Password: password,
		DB:       db,
	})

	//	Rdb.Set(RdbContext,"key","value",time.Minute*30) 添加key-value 存活时间为30min
	//	val, err := Rdb.Get(RdbContext, "key").Result()
	//	if err == redis.Nil {
	//		fmt.Println("key don't exist")
	//	}
	//	fmt.Println(val)
}
