package main

import (
	"fmt"
	"github.com/itziklavon/kit2go/tree/master/redis-helper"
	"github.com/garyburd/redigo/redis"
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
