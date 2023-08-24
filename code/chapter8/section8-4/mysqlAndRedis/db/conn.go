package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

// Connect 创建连接MySQL的句柄
func Connect(conf Conf) (err error) {
	if Conn != nil {
		return
	}

	dsn := fillConnArgs(conf)

	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用表名复数
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	err = ping()
	if err != nil {
		return err
	}
	
	return nil
}

// fillConnArgs 根据配置拼接连接数据库的必要信息
func fillConnArgs(conf Conf) (args string) {
	return conf.User + ":" + conf.Password + "@tcp(" + conf.Domain +
		":" + conf.Port + ")/" + conf.Name + "?charset=utf8&parseTime=True&loc=Local"
}

// ping 测试数据库连接是否正常
func ping() (err error) {
	sqlDB, err := Conn.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	return nil
}
