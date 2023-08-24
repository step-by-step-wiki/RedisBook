package main

import (
	"github.com/gin-gonic/gin"
	"mysqlAndRedis/cache"
	"mysqlAndRedis/controller"
	"mysqlAndRedis/db"
)

func main() {
	mysqlConf := db.Conf{
		Domain:   "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "123456",
		Name:     "redisStudy",
	}
	err := db.Connect(mysqlConf)
	if err != nil {
		panic("connect mysql failed:" + err.Error())
	}

	redisConf := cache.Conf{
		NetWork:  "tcp",
		Address:  "localhost:6379",
		User:     "redis_user",
		Password: "redis_password",
	}

	err = cache.Connect(redisConf)
	if err != nil {
		panic("connect redis failed:" + err.Error())
	}

	// 进程结束前关闭连接
	sqlDB, err := db.Conn.DB()
	defer sqlDB.Close()
	defer cache.Conn.Close()

	// 确认连接成功后开启web服务
	r := gin.Default()
	r.POST("/student/getById", controller.GetStudentById)
	r.Run("0.0.0.0:8085")
}
