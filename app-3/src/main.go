package main

import (
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/itziklavon/kit2go/http-client-helper/src/http_client_helper"
	"github.com/itziklavon/kit2go/redis-helper/src/redis_helper"
)

var done = make(chan bool)

func main() {
	brandSet := redis_helper.GetBrandSet()
	fmt.Println(brandSet)
	for _, brandId := range brandSet {
		fmt.Println("staring for brand:", brandId)
		go handleMessges(brandId)
	}
	fmt.Println(<-done)
}

func handleMessges(brandId int) {
	for {
		redisHelper := redis_helper.GetRedisConnection(brandId)
		c, psc := redisHelper.Subscribe("__keyevent@0__:expired")
		for c.Err() == nil {
			switch v := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
				fmt.Println()
				messageData := string(v.Data[:])

				fmt.Println("message is: " + messageData[:len(messageData)])
				if strings.HasPrefix(messageData, "TKN_") {
					logoutPlayer(brandId, messageData[4:len(messageData)])
				}
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
				fmt.Println()
			case error:
				fmt.Println(c.Err())
				fmt.Println()
			}
		}
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
