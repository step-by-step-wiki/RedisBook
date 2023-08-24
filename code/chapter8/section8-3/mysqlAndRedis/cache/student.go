package cache

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
)

const StudentKeyPrefix = "Stu"

const ExpireSecond = 3600

type Student struct {
	Id    int
	Name  string
	Age   int
	Score float64
	Exist bool
}

func (s *Student) FindById(id int) (err error) {
	// 判断键是否存在
	s.Exist, err = s.exists(id)
	if err != nil {
		return err
	}

	// 键不存在则直接返回
	if !s.Exist {
		return nil
	}

	// 确认键存在,则从redis中读取(有可能读到的是一个0值)
	idStr := strconv.Itoa(id)
	key := StudentKeyPrefix + idStr
	reply, err := redis.Values(Conn.Do("LRANGE", key, 0, -1))
	if err != nil {
		return err
	}

	s.Id, err = redis.Int(reply[0], nil)
	if err != nil {
		return err
	}

	s.Name, err = redis.String(reply[1], nil)
	if err != nil {
		return err
	}

	s.Age, err = redis.Int(reply[2], nil)
	if err != nil {
		return err
	}

	s.Score, err = redis.Float64(reply[3], nil)
	if err != nil {
		return err
	}

	// 若读取了该key 则重置过期时间
	// 此处忽略错误
	_ = s.setExpireTime(id)

	return nil
}

func (s *Student) SaveById(id int) error {
	idStr := strconv.Itoa(id)
	key := StudentKeyPrefix + idStr
	_, err := Conn.Do("RPUSH", key, s.Id, s.Name, s.Age, s.Score)
	if err != nil {
		return err
	}

	err = s.setExpireTime(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Student) exists(id int) (bool, error) {
	idStr := strconv.Itoa(id)
	key := StudentKeyPrefix + idStr
	reply, err := redis.Int(Conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return reply == 1, nil
}

func (s *Student) setExpireTime(id int) error {
	idStr := strconv.Itoa(id)
	key := StudentKeyPrefix + idStr
	_, err := Conn.Do("EXPIRE", key, ExpireSecond)
	return err
}
