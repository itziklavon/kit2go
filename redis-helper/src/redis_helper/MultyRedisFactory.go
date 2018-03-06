package redis_helper

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/itziklavon/kit2go/configuration/src/configuration"
	"github.com/itziklavon/kit2go/http-client-helper/src/http_client_helper"
	"github.com/itziklavon/kit2go/general-log/src/general_log"
)

var DISCOVERY_URL = configuration.GetPropertyValue("DISCOVERY_URL")

var redisConnections map[int]string

type RedisData struct {
	BrandId int    `json:"brandId"`
	Uri     string `json:"uri"`
}

func GetRedisConnection(brandId int) *redis.Client {
	if len(redisConnections) == 0 {
		initMap()
	}
	return getBrandRedisConnection(redisConnections[brandId])
}

func initMap() {
	redisConnections = make(map[int]string)
	url := DISCOVERY_URL + "discovery-web/brand/services/REDIS"
	httpResponse := http_client_helper.GETBody(url, nil)
	var arr []RedisData
	_ = json.Unmarshal([]byte(httpResponse), &arr)
	for i := 0; i < len(arr); i = i + 1 {
		redisConnections[arr[i].BrandId] = arr[i].Uri
	}
}

func getBrandRedisConnection(host string) *redis.Client {
	return getRedisConnectionByHost(host)
}

func KeysMuyltiBrand(brandId int) []string {
	conn := GetRedisConnection(brandId)
	value, err := conn.Keys("*").Result()
	if err != nil {
		general_log.ErrorException(":Keys: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func KeysWithPatternuyltiBrand(brandId int, pattern string) []string {
	conn := GetRedisConnection(brandId)
	value, err := conn.Keys(pattern).Result()
	if err != nil {
		general_log.ErrorException(":Keys: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func GetuyltiBrand(brandId int, key string) string {
	conn := GetRedisConnection(brandId)
	value, err := conn.Get(key).Result()
	if err != nil {
		general_log.ErrorException(":Get: couldn't get key from redis: "+key, err)
	}
	defer conn.Close()
	return value
}

func HGetuyltiBrand(brandId int, key string, hkey string) string {
	conn := GetRedisConnection(brandId)
	value, err := conn.HGet(key, hkey).Result()
	if err != nil {
		general_log.ErrorException(":Set: couldn't get key from redis: "+key+", hKey: "+hkey, err)
	}
	defer conn.Close()
	return value
}

func GetSysParamuyltiBrand(brandId int, hkey string) string {
	return HGetuyltiBrand(brandId, "SysParams", hkey)
}

func GetBrandIduyltiBrand(brandId int) string {
	return GetSysParamuyltiBrand(brandId, "GS_BRAND_ID")
}
