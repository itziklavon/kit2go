package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
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
