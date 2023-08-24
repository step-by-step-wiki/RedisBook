package cache

// Conf redis相关配置
type Conf struct {
	// NetWork 网络类型
	NetWork string
	// Address Redis地址 格式: ip:port
	Address string
	// User 用户名
	User string
	// Password 密码
	Password string
}
