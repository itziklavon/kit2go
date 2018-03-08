package redis_helper

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/itziklavon/kit2go/configuration/src/configuration"
	"github.com/itziklavon/kit2go/general-log/src/general_log"
	"github.com/itziklavon/kit2go/http-client-helper/src/http_client_helper"
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

func KeysMultyBrand(brandId int) []string {
	conn := GetRedisConnection(brandId)
	value, err := conn.Keys("*").Result()
	if err != nil {
		general_log.ErrorException(":KeysMultyBrand: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func KeysWithPatternMultyBrand(brandId int, pattern string) []string {
	conn := GetRedisConnection(brandId)
	value, err := conn.Keys(pattern).Result()
	if err != nil {
		general_log.ErrorException(":KeysWithPatternMultyBrand: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func GetMultyBrand(brandId int, key string) string {
	conn := GetRedisConnection(brandId)
	value, err := conn.Get(key).Result()
	if err != nil {
		general_log.ErrorException(":GetMultyBrand: couldn't get key from redis: "+key, err)
	}
	defer conn.Close()
	return value
}

func HGetMultyBrand(brandId int, key string, hkey string) string {
	conn := GetRedisConnection(brandId)
	value, err := conn.HGet(key, hkey).Result()
	if err != nil {
		general_log.ErrorException(":HGetMultyBrand: couldn't get key from redis: "+key+", hKey: "+hkey, err)
	}
	defer conn.Close()
	return value
}

func HGetAllMultyBrand(brandId int, key string) map[string]string {
	conn := GetRedisConnection(brandId)
	value, err := conn.HGetAll(key).Result()
	if err != nil {
		general_log.ErrorException(":HGetAllMultyBrandSet: couldn't get key from redis: "+key, err)
	}
	defer conn.Close()
	return value
}

func SetMultyBrand(brandId int, key string, value string) {
	conn := GetRedisConnection(brandId)
	err := conn.Set(key, value, 0).Err()
	if err != nil {
		general_log.ErrorException(":SetMultyBrand: couldn't get key from redis: "+key, err)
	}
	defer conn.Close()
}

func HSetMultyBrand(brandId int, key string, hkey string, value string) {
	conn := GetRedisConnection(brandId)
	err := conn.HSet(key, hkey, value).Err()
	if err != nil {
		general_log.ErrorException(":HSetMultyBrand: couldn't get key from redis: "+key+", hKey: "+hkey, err)
	}
	defer conn.Close()
}
func PubSubscribe(brandId int, pattern string) *redis.PubSub {
	conn := GetRedisConnection(brandId)
	pSubscribe := conn.PSubscribe(pattern)
	defer conn.Close()
	return pSubscribe
}

func GetSysParamMultyBrand(brandId int, hkey string) string {
	return HGetMultyBrand(brandId, "SysParams", hkey)
}

func GetBrandIdMultyBrand(brandId int) string {
	return GetSysParamMultyBrand(brandId, "GS_BRAND_ID")
}
