package redis

import "github.com/go-redis/redis/v8"

var RD *redis.Client
func init()  {
	RD = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
}
