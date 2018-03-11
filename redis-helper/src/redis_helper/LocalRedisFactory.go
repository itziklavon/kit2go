package redis_helper

import (
	"github.com/itziklavon/kit2go/configuration/src/configuration"
)

var redisHost = configuration.GetPropertyValue("REDIS_HOST")

func GetRedisLocalConnection() RedisSessionHelper {
	return RedisSessionHelper{host: redisHost}
}
