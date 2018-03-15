package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../configuration/src/configuration"
	"../../general-log/src/general_log"
	"../../http-client-helper/src/http_client_helper"
	"../../redis-helper/src/redis_helper"
	"github.com/garyburd/redigo/redis"
)

var fileName = configuration.GetLogPropertyValue("REDIS_LISTENER_LOG")

var done = make(chan bool)

func main() {
	general_log.SetLogOutput(fileName)
	brandSet := redis_helper.GetBrandSet()
	general_log.Debug(":main: extracted brands are: ", brandSet)
	for _, brandId := range brandSet {
		general_log.Debug(":main: staring for brand: " + string(brandId))
		go handleMessges(brandId)
	}
	general_log.Debug(strconv.FormatBool(<-done))
}

func handleMessges(brandId int) {
	for {
		redisHelper := redis_helper.GetRedisConnection(brandId)
		c, psc := redisHelper.Subscribe("__keyevent@0__:expired")
		for c.Err() == nil {
			switch v := psc.Receive().(type) {
			case redis.Message:
				messageData := string(v.Data[:])
				general_log.Debug(":handleMessges: exracted message is: " + messageData)
				if strings.HasPrefix(messageData, "TKN_") {
					if redisHelper.Exists(messageData[4:len(messageData)]) {
						go logoutPlayer(brandId, messageData[4:len(messageData)])
					}
				}
			}
		}
		general_log.ErrorException(":handleMessges: an error occured in brand: "+string(brandId), c.Err())
		c.Close()
	}
	done <- true
}

func logoutPlayer(brandId int, token string) {
	url := http_client_helper.GetDiscoveryUrl(brandId, "GSS") + "/player/logout"
	values := map[string]string{"auth_token": token}
	fmt.Println("sending message to brand 7 with token:" + token)

	http_client_helper.POST(url, values, nil)
}
