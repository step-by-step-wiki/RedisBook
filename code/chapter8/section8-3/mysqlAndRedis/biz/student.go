package biz

import (
	"errors"
	"gorm.io/gorm"
	"mysqlAndRedis/cache"
	"mysqlAndRedis/db"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score float64
}

func (s *Student) GetById(id int) error {
	// 从缓存中获取
	cacheStu := &cache.Student{}
	err := cacheStu.FindById(id)
	if err != nil {
		return err
	}

	// 如果缓存中存在,则直接返回(虽然有可能返回零值)
	if cacheStu.Exist {
		s.Id = cacheStu.Id
		s.Name = cacheStu.Name
		s.Age = cacheStu.Age
		s.Score = cacheStu.Score
		return nil
	}

	// 从数据库中获取
	dbStu := &db.Student{}
	err = dbStu.FindById(id)
	// 如果数据库中不存在 则同样将对应的键写入Redis(此时存储的是一个零值)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	s.Id = dbStu.Id
	s.Name = dbStu.Name
	s.Age = dbStu.Age
	s.Score = dbStu.Score

	// 写入缓存
	cacheStu.Id = dbStu.Id
	cacheStu.Name = dbStu.Name
	cacheStu.Age = dbStu.Age
	cacheStu.Score = dbStu.Score
	err = cacheStu.SaveById(id)
	if err != nil {
		return err
	}

	return nil
}
