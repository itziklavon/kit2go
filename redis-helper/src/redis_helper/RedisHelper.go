package redis_helper

import (
	"github.com/go-redis/redis"
	"github.com/itziklavon/kit2go/general-log/src/general_log"
)

type RedisSessionHelper struct {
	host string
}

func (r RedisSessionHelper) getRedisConnection() *redis.Client {
	return getRedisConnectionByHost(r.host)
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

func (r RedisSessionHelper) Keys() []string {
	conn := r.getRedisConnection()
	value, err := conn.Keys("*").Result()
	if err != nil {
		general_log.ErrorException(":Keys: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func (r RedisSessionHelper) KeysWithPattern(pattern string) []string {
	conn := r.getRedisConnection()
	value, err := conn.Keys(pattern).Result()
	if err != nil {
		general_log.ErrorException(":Keys: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func (r RedisSessionHelper) Get(key string) string {
	conn := r.getRedisConnection()
	value, err := conn.Get(key).Result()
	if err != nil {
		general_log.ErrorException(":Get: couldn't get key from redis: "+key, err)
	}
	defer conn.Close()
	return value
}

func (r RedisSessionHelper) HGet(key string, hkey string) string {
	conn := r.getRedisConnection()
	value, err := conn.HGet(key, hkey).Result()
	if err != nil {
		general_log.ErrorException(":Set: couldn't get key from redis: "+key+", hKey: "+hkey, err)
	}
	defer conn.Close()
	return value
}

func (r RedisSessionHelper) GetSysParam(hkey string) string {
	return r.HGet("SysParams", hkey)
}

func (r RedisSessionHelper) GetBrandId() string {
	return r.GetSysParam("GS_BRAND_ID")
}
