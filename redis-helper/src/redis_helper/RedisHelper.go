package redis_helper

import (
	"github.com/go-redis/redis"
	"github.com/itziklavon/kit2go/general-log/src/general_log"
)

var redisHost = configuration.GetPropertyValue("REDIS_HOST")

func getRedisConnection() *redis.Client {
	return getRedisConnectionByHost(redisHost)
}

func getRedisConnectionByHost(host string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		general_log.ErrorException(":getRedisConnection: couldn't connect ro redis", err)
	}
	return client
}

func Keys() []string {
	conn := getRedisConnection()
	value, err := conn.Keys("*").Result()
	if err != nil {
		general_log.ErrorException(":Keys: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func KeysWithPattern(pattern string) []string {
	conn := getRedisConnection()
	value, err := conn.Keys(pattern).Result()
	if err != nil {
		general_log.ErrorException(":Keys: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func Get(key string) string {
	conn := getRedisConnection()
	value, err := conn.Get(key).Result()
	if err != nil {
		general_log.ErrorException(":Get: couldn't get key from redis: "+key, err)
	}
	defer conn.Close()
	return value
}

func HGet(key string, hkey string) string {
	conn := getRedisConnection()
	value, err := conn.HGet(key, hkey).Result()
	if err != nil {
		general_log.ErrorException(":Set: couldn't get key from redis: "+key+", hKey: "+hkey, err)
	}
	defer conn.Close()
	return value
}

func GetSysParam(hkey string) string {
	return HGet("SysParams", hkey)
}

func GetBrandId() string {
	return GetSysParam("GS_BRAND_ID")
}
