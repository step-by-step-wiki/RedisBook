package cache

import (
	"github.com/gomodule/redigo/redis"
)

var Conn redis.Conn

// Connect 根据给定的配置 创建Redis连接
func Connect(conf Conf) (err error) {
	if Conn != nil {
		return nil
	}

	// 连接到Redis
	Conn, err = redis.Dial(conf.NetWork, conf.Address)
	if err != nil {
		return err
	}

	if conf.User != "" || conf.Password != "" {
		err = auth(conf, Conn)
		if err != nil {
			return err
		}
	}

	err = ping(Conn)
	if err != nil {
		return err
	}

	return nil
}

// auth 根据给定的配置和连接 进行认证
func auth(conf Conf, conn redis.Conn) (err error) {
	_, err = conn.Do("AUTH", conf.User, conf.Password)
	return err
}

// ping 根据给定的连接 进行探活
func ping(conn redis.Conn) (err error) {
	_, err = conn.Do("PING")
	return err
}
