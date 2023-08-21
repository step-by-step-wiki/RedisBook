package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"operateRedis/conn"
	"operateRedis/operate"
	"time"
)

func main() {
	conf := conn.Conf{
		NetWork:  "tcp",
		Address:  "localhost:6379",
		User:     "redis_user",
		Password: "redis_password",
	}

	poolConf := conn.PoolConf{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: time.Hour,
		Conf:        conf,
	}

	// 初始化连接池
	conn.NewPool(poolConf)

	// 从连接池中获取一个连接
	redisConn, err := conn.GetConnFromPool(poolConf)
	if err != nil {
		log.Fatalf("get conn from pool error: %v\n", err)
	}

	// 将该连接返回到连接池中
	defer conn.CloseConnToPool(redisConn)

	// 使用该连接进行操作
	operateNum, err := operate.Pipeline(redisConn)
	if err != nil {
		log.Fatalf("pipeline error: %v\n", err)
	}

	// 读取响应
	for i := 0; i < operateNum; i++ {
		reply, err := redis.String(redisConn.Receive())
		if err != nil {
			log.Fatalf("receive error: %v\n", err)
		}

		log.Printf("reply: %v\n", reply)
	}
}
