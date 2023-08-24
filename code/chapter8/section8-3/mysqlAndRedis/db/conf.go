package db

// Conf 数据库相关配置
type Conf struct {
	// Domain 数据库服务器IP地址
	Domain string
	// Port 数据库端口
	Port string
	// User 用户名
	User string
	// Password 密码
	Password string
	// Name 数据库名
	Name string
}
