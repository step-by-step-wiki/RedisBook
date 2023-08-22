package list

import (
	"github.com/gomodule/redigo/redis"
)

func LPush(conn redis.Conn, key string, values ...string) (int, error) {
	length := 0

	for _, value := range values {
		err := conn.Send("LPUSH", key, value)
		if err != nil {
			return 0, err
		}
	}
	err := conn.Flush()
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(values); i++ {
		reply, _ := redis.Int(conn.Receive())
		length = reply
	}

	return length, nil
}

func RPush(conn redis.Conn, key string, values ...string) (int, error) {
	length := 0

	for _, value := range values {
		err := conn.Send("RPUSH", key, value)
		if err != nil {
			return 0, err
		}
	}
	err := conn.Flush()
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(values); i++ {
		reply, _ := redis.Int(conn.Receive())
		length = reply
	}

	return length, nil
}

func LPop(conn redis.Conn, key string) (string, error) {
	return redis.String(conn.Do("LPOP", key))
}

func RPop(conn redis.Conn, key string) (string, error) {
	return redis.String(conn.Do("RPOP", key))
}

func LLen(conn redis.Conn, key string) (int, error) {
	return redis.Int(conn.Do("LLEN", key))
}

func LRange(conn redis.Conn, key string, start int, end int) ([]string, error) {
	return redis.Strings(conn.Do("LRANGE", key, start, end))
}

func LTrim(conn redis.Conn, key string, start int, end int) (string, error) {
	return redis.String(conn.Do("LTRIM", key, start, end))
}

func LSet(conn redis.Conn, key string, index int, value string) (string, error) {
	return redis.String(conn.Do("LSET", key, index, value))
}

func LIndex(conn redis.Conn, key string, index int) (string, error) {
	return redis.String(conn.Do("LINDEX", key, index))
}
