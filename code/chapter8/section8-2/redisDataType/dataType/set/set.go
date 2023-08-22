package set

import "github.com/gomodule/redigo/redis"

func SAdd(conn redis.Conn, key string, members ...interface{}) (int, error) {
	args := redis.Args{}.Add(key).AddFlat(members)
	return redis.Int(conn.Do("SADD", args...))
}

func SIsMember(conn redis.Conn, key string, member interface{}) (bool, error) {
	return redis.Bool(conn.Do("SISMEMBER", key, member))
}

func SMembers(conn redis.Conn, key string) ([]string, error) {
	return redis.Strings(conn.Do("SMEMBERS", key))
}

// SRem 删除集合中的元素(根据给定的元素值删除) 返回删除的元素个数
func SRem(conn redis.Conn, key string, members ...interface{}) (int, error) {
	args := redis.Args{}.Add(key).AddFlat(members)
	return redis.Int(conn.Do("SREM", args...))
}

// SPop 随机删除并返回集合中的一个元素
func SPop(conn redis.Conn, key string) (string, error) {
	return redis.String(conn.Do("SPOP", key))
}
