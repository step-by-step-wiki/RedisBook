package operate

import "github.com/gomodule/redigo/redis"

func Get(conn redis.Conn, key string) (reply string, err error) {
	return redis.String(conn.Do("GET", key))
}
