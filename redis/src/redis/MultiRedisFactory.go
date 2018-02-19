package redis


import (
	"github.com/itziklavon/kit2go/general-log/src/general_log"

	"github.com/itziklavon/kit2go/http-client/src/http"

	"github.com/itziklavon/kit2go/configuration/src/configuration"
	"menteslibres.net/gosexy/redis"

	"encoding/json"
)

var DISCOVERY_URL = configuration.GetPropertyValue("DISCOVERY_URL")


var redisConnections map[int]String

type RedisData struct {
	BrandId int `json:"brandId"`
	Uri string `json:"uri"`
}

func GetRedisConnection(brandId int) string {
	if len(redisConnections) == 0 {
		initMap()
	}
	return redisConnections[brandId]
}

func initMap() {
	redisConnections = make(map[int]string)
	url := DISCOVERY_URL + "discovery-web/brand/services/REDIS"
	httpResponse := http.GETBody(url, nil)
	var arr []RedisData
	_ = json.Unmarshal([]byte(httpResponse), &arr)
	for i := 0; i < len(arr); i = i + 1 {
		redisConnections[arr[i].BrandId] = arr[i].Uri
	}
}



