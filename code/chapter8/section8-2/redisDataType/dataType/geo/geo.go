package geo

import (
	"github.com/gomodule/redigo/redis"
)

func GeoAdd(conn redis.Conn, key string, geoMap map[string][]float64) (int, error) {
	args := redis.Args{}.Add(key)
	for locationName, geo := range geoMap {
		args = args.AddFlat(geo).AddFlat(locationName)
	}

	return redis.Int(conn.Do("GEOADD", args...))
}

func GeoPos(conn redis.Conn, key string, locationNames ...string) ([]*[2]float64, error) {
	args := redis.Args{}.Add(key).AddFlat(locationNames)
	return redis.Positions(conn.Do("GEOPOS", args...))
}

func GeoDist(conn redis.Conn, key string, locationName1, locationName2, unit string) (float64, error) {
	return redis.Float64(conn.Do("GEODIST", key, locationName1, locationName2, unit))
}

func GeoRadius(conn redis.Conn, key string, longitude, latitude, radius float64, unit string, withCoord, withDist, withHash bool, count int, sort string) ([]interface{}, error) {
	args := redis.Args{}.Add(key).Add(longitude).Add(latitude).Add(radius).Add(unit)
	if withCoord {
		args = args.Add("WITHCOORD")
	}
	if withDist {
		args = args.Add("WITHDIST")
	}
	if withHash {
		args = args.Add("WITHHASH")
	}
	if count != 0 {
		args = args.Add("COUNT", count)
	}
	if sort != "" {
		args = args.Add(sort)
	}

	return redis.Values(conn.Do("GEORADIUS", args...))
}
