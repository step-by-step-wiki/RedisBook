package cache

import "mysqlAndRedis/lib"

const ExpireSecond = 3600

const CeilSecond = 60

func genExpireSecond() int {
	return ExpireSecond + lib.GenRandInt(CeilSecond)
}
