package lock

import (
	"github.com/gomodule/redigo/redis"
	"math/rand"
)

type DistributeLock struct {
	Key   string
	value int
	TTL   int
	Conn  redis.Conn
}

// Acquire 获取锁
func (l *DistributeLock) Acquire() bool {
	result, err := redis.String(l.Conn.Do("SET", l.Key, l.value, "NX", "EX", l.TTL))
	if err != nil {
		return false
	}

	return result == "OK"
}

// Release 释放锁
func (l *DistributeLock) Release() bool {
	// KEY存在 且 VALUE和设置时的值相同,才能删除
	luaScript := `
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return 0
end
`
	script := redis.NewScript(1, luaScript)
	result, err := redis.Int(script.Do(l.Conn, l.Key, l.value))
	if err != nil {
		return false
	}

	return result == 1
}

// GenValue 生成随机值作为锁的value
func (l *DistributeLock) GenValue() {
	l.value = rand.Int()
}
