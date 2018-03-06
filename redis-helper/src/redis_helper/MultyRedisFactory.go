package redis_helper

import (
	"encoding/json"
	"github.com/itziklavon/kit2go/configuration/src/configuration"
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
	return redisConnections[brandId]
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
