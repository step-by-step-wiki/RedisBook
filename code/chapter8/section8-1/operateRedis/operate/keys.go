package operate

import "github.com/gomodule/redigo/redis"

func Keys(conn redis.Conn, pattern string) (reply []string, err error) {
	return redis.Strings(conn.Do("KEYS", pattern))
}
