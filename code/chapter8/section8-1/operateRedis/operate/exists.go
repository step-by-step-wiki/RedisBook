package operate

import "github.com/gomodule/redigo/redis"

func Exists(conn redis.Conn, key string) (reply bool, err error) {
	return redis.Bool(conn.Do("EXISTS", key))
}
