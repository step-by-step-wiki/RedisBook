package operate

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

func Transaction(conn redis.Conn) (replies []string, err error) {
	// 1. Watch
	_, err = redis.String(conn.Do("WATCH", "counter"))
	if err != nil {
		return nil, err
	}

	// 若counter的值大于10 则不执行事务
	// 此处假定10为counter的原值
	counter, _ := redis.Int(conn.Do("GET", "counter"))
	if counter > 10 {
		conn.Do("UNWATCH")
		return nil, errors.New("counter value is too high. Not continuing the transaction")
	}

	// 2. Multi
	err = conn.Send("MULTI")
	if err != nil {
		return nil, err
	}

	// 3. Exec
	err = conn.Send("SET", "transaction_key_1", "transaction_value_1")
	if err != nil {
		conn.Do("DISCARD")
		return nil, err
	}

	err = conn.Send("SET", "transaction_key_2", "transaction_value_2")
	if err != nil {
		conn.Do("DISCARD")
		return nil, err
	}

	replies, err = redis.Strings(conn.Do("EXEC"))
	if err != nil {
		conn.Do("DISCARD")
		return nil, err
	}

	return replies, nil
}
