package operate

import "github.com/gomodule/redigo/redis"

func Set(conn redis.Conn, key string, value string) (reply string, err error) {
	return redis.String(conn.Do("SET", key, value))
}
