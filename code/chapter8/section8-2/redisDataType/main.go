package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"redisDataType/conn"
	"redisDataType/dataType/geo"
	"time"
)

func main() {
	conf := conn.Conf{
		NetWork:  "tcp",
		Address:  "localhost:6379",
		User:     "redis_user",
		Password: "redis_password",
	}

	poolConf := conn.PoolConf{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: time.Hour,
		Conf:        conf,
	}

	// 初始化连接池
	conn.NewPool(poolConf)

	// 从连接池中获取一个连接
	redisConn, err := conn.GetConnFromPool(poolConf)
	if err != nil {
		log.Fatalf("get conn from pool error: %v\n", err)
	}

	// 将该连接返回到连接池中
	defer conn.CloseConnToPool(redisConn)

	// 添加地理位置
	locations := map[string][]float64{
		"location1": {13.3615, 38.1157},
		"location2": {15.0873, 37.5027},
	}
	addNum, err := geo.GeoAdd(redisConn, "myLocations", locations)
	if err != nil {
		log.Fatalf("geo add error: %v\n", err)
	}
	log.Printf("add %d locations\n", addNum)

	locations = map[string][]float64{
		"location2": {16.3615, 41.1157},
		"location3": {14.3615, 39.1157},
	}
	addNum, err = geo.GeoAdd(redisConn, "myLocations", locations)
	if err != nil {
		log.Fatalf("geo add error: %v\n", err)
	}
	log.Printf("add %d locations\n", addNum)

	// 获取指定位置的经纬度
	locationNames := []string{"location1", "location2", "location3"}
	positions, err := geo.GeoPos(redisConn, "myLocations", locationNames...)
	if err != nil {
		log.Fatalf("geo pos error: %v\n", err)
	}
	for i, position := range positions {
		log.Printf("location: %s, longitude: %f, latitude: %f\n", locationNames[i], position[0], position[1])
	}

	// 计算2个位置之间的距离
	distance, err := geo.GeoDist(redisConn, "myLocations", "location1", "location2", "km")
	if err != nil {
		log.Fatalf("geo dist error: %v\n", err)
	}
	log.Printf("distance between location1 and location2 is %.2f km\n", distance)

	// 获取指定位置半径范围内的位置
	radiusLocationValues, err := geo.GeoRadius(redisConn, "myLocations", 15, 37, 200, "km", true, true, false, 0, "ASC")
	if err != nil {
		log.Fatalf("geo radius error: %v\n", err)
	}

	for _, locationValue := range radiusLocationValues {
		location, _ := redis.Strings(locationValue, nil)

		// 单独处理最后一个元素 因为最后一个元素表示经纬度 是一个slice
		positionV, _ := redis.Values(locationValue, nil)
		position, _ := redis.Float64s(positionV[len(positionV)-1], nil)

		log.Printf("location: %v, distance: %v, position: %v \n", location[0], location[1], position)
	}
}
