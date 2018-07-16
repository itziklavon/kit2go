package redis_helper

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/itziklavon/kit2go/general-log/src/general_log"
)

type RedisSessionHelper struct {
	host string
}

func (r RedisSessionHelper) getRedisConnection() *redis.Pool {
	return newPool(r.host + ":6379")
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func getRedisConnectionByHost(host string) redis.Conn {
	general_log.Debug(":getRedisConnectionByHost: initializing host with param:" + host)
	redisHelper, err := redis.Dial("tcp", host)
	if err != nil {
		general_log.ErrorException(":getRedisConnectionByHost: couldn't connect ro redis", err)
	}
	return redisHelper
}

func (r RedisSessionHelper) Keys() []string {
	pool := r.getRedisConnection()
	conn := pool.Get()
	defer conn.Close()
	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", "*"))
		if err != nil {
			general_log.ErrorException(":Keys: an error occurred", err)
			return keys
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}
	return keys
}

func (r RedisSessionHelper) Get(key string) string {
	pool := r.getRedisConnection()
	conn := pool.Get()
	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		general_log.ErrorException(":Get: an error occurred", err)
		return string(data)
	}
	defer conn.Close()
	return string(data)
}

func (r RedisSessionHelper) HGet(key string, hkey string) string {
	pool := r.getRedisConnection()
	conn := pool.Get()
	var data []byte
	data, err := redis.Bytes(conn.Do("HGET", key, hkey))
	if err != nil {
		general_log.ErrorException(":HGet: an error occurred", err)
		return string(data)
	}
	defer conn.Close()
	return string(data)
}

func (r RedisSessionHelper) Exists(key string) bool {
	pool := r.getRedisConnection()
	conn := pool.Get()
	defer conn.Close()
	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		general_log.ErrorException(":Exists: an error occurred", err)
		return ok
	}
	return ok
}

func (r RedisSessionHelper) GetSysParam(hkey string) string {
	return r.HGet("SysParams", hkey)
}

func (r RedisSessionHelper) GetBrandId() string {
	return r.GetSysParam("GS_BRAND_ID")
}

func (r RedisSessionHelper) Subscribe(channel string) (redis.Conn, redis.PubSubConn) {
	pool := r.getRedisConnection()
	conn := pool.Get()
	psc := redis.PubSubConn{conn}
	psc.Subscribe(channel)
	return conn, psc
}

func (r RedisSessionHelper) ConfigSet(key string, value string) {
	pool := r.getRedisConnection()
	conn := pool.Get()
	conn.Do("CONFIG", "SET", "notify-keyspace-events", "KEA")
}
