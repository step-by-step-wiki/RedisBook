package operate

import "github.com/gomodule/redigo/redis"

func Del(conn redis.Conn, key string) (reply int, err error) {
	return redis.Int(conn.Do("DEL", key))
}
