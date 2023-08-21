package conn

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

type Conf struct {
	NetWork  string // NetWork 网络类型
	Address  string // Address Redis地址 格式: ip:port
	User     string // User 用户名
	Password string // Password 密码
}

// NewConn 根据给定的配置 创建Redis连接
func NewConn(conf *Conf) (conn redis.Conn, err error) {
	// 连接到Redis
	conn, err = redis.Dial(conf.NetWork, conf.Address)
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
		return
	}

	if conf.User != "" || conf.Password != "" {
		err = auth(conf, conn)
		if err != nil {
			log.Fatalf("Could not authenticate: %v\n", err)
			return
		}
	}

	err = ping(conn)
	if err != nil {
		log.Fatalf("Could not ping: %v\n", err)
		return
	}

	return
}

// auth 根据给定的配置和连接 进行认证
func auth(conf *Conf, conn redis.Conn) (err error) {
	_, err = conn.Do("AUTH", conf.User, conf.Password)
	return err
}

// ping 根据给定的连接 进行探活
func ping(conn redis.Conn) (err error) {
	_, err = conn.Do("PING")
	return err
}
