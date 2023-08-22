package sortedSet

import (
	"github.com/gomodule/redigo/redis"
)

func ZAdd(conn redis.Conn, key string, scoreMap map[int]interface{}) (int, error) {
	args := redis.Args{}.Add(key).AddFlat(scoreMap)
	return redis.Int(conn.Do("ZADD", args...))
}

func ZCard(conn redis.Conn, key string) (int, error) {
	return redis.Int(conn.Do("ZCARD", key))
}

// ZRange 根据索引获取有序集合中的元素(按照分数从小到大排序)
func ZRange(conn redis.Conn, key string, start int, stop int) ([]string, error) {
	return redis.Strings(conn.Do("ZRANGE", key, start, stop))
}

// ZRevRange 根据索引获取有序集合中的元素(按照分数从大到小排序)
func ZRevRange(conn redis.Conn, key string, start int, stop int) ([]string, error) {
	return redis.Strings(conn.Do("ZREVRANGE", key, start, stop))
}

func ZScore(conn redis.Conn, key string, member string) (float64, error) {
	return redis.Float64(conn.Do("ZSCORE", key, member))
}

func ZRem(conn redis.Conn, key string, members ...string) (int, error) {
	args := redis.Args{}.Add(key).AddFlat(members)
	return redis.Int(conn.Do("ZREM", args...))
}
