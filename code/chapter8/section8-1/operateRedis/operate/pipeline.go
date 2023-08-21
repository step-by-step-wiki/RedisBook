package operate

import "github.com/gomodule/redigo/redis"

// Pipeline 以管道方式一次性发送多个命令
func Pipeline(conn redis.Conn) (operateNum int, err error) {
	operateNum = 0
	err = conn.Send("SET", "name", "Peter")
	if err != nil {
		return 0, err
	}
	operateNum++

	err = conn.Send("GET", "name")
	if err != nil {
		return 0, err
	}
	operateNum++

	err = conn.Send("SET", "age", "18")
	if err != nil {
		return 0, err
	}
	operateNum++

	err = conn.Send("GET", "age")
	if err != nil {
		return 0, err
	}
	operateNum++

	// 将缓冲区的命令写入到连接中
	err = conn.Flush()
	if err != nil {
		return 0, err
	}

	return operateNum, nil
}
