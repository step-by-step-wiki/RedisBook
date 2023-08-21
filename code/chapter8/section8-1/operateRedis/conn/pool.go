package conn

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type PoolConf struct {
	// MaxIdle 最大空闲连接数
	MaxIdle int
	// MaxActive 最大连接数，0表示无限制
	MaxActive int
	// IdleTimeout 空闲连接超时时间
	IdleTimeout time.Duration
	// Conf Redis连接配置
	Conf Conf
}

var Pool *redis.Pool

func NewPool(poolConf PoolConf) {
	Pool = &redis.Pool{
		MaxIdle:     poolConf.MaxIdle,
		MaxActive:   poolConf.MaxActive,
		IdleTimeout: poolConf.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(poolConf.Conf.NetWork, poolConf.Conf.Address)
		},
	}
}

// GetConnFromPool 从连接池中获取一个经过认证后的连接
func GetConnFromPool(poolConf PoolConf) (conn redis.Conn, err error) {
	conn = Pool.Get()
	err = auth(&poolConf.Conf, conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// CloseConnToPool 将该连接返回到连接池中
func CloseConnToPool(conn redis.Conn) error {
	// 关闭一个从连接池中取出的连接时,使用conn.Close()方法
	// 并不会让该连接真正的关闭,而是将该连接返回到连接池中,以便后续复用
	return conn.Close()
}
