package hash

import (
	"github.com/gomodule/redigo/redis"
)

func HSet(conn redis.Conn, key string, hashes map[string]string) (int, error) {
	_, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(hashes)...)
	if err != nil {
		return 0, err
	}

	length, err := HLen(conn, key)
	if err != nil {
		// TODO: 此处获取长度失败不应该返回0
		return 0, err
	}

	return length, nil
}

func HLen(conn redis.Conn, key string) (int, error) {
	return redis.Int(conn.Do("HLEN", key))
}

func HGet(conn redis.Conn, key string, field string) (string, error) {
	return redis.String(conn.Do("HGET", key, field))
}

func HGetAll(conn redis.Conn, key string) (map[string]string, error) {
	return redis.StringMap(conn.Do("HGETALL", key))
}

func HExists(conn redis.Conn, key string, field string) (bool, error) {
	return redis.Bool(conn.Do("HEXISTS", key, field))
}

func HDel(conn redis.Conn, key string, fields ...string) (int, error) {
	args := redis.Args{}.Add(key).AddFlat(fields)
	return redis.Int(conn.Do("HDEL", args...))
}
