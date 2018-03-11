package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	getMessages()
}

func getMessages() {

	for {
		// Get a connection from a pool
		c, err := redis.Dial("tcp", "172.17.30.17:6379")
		if err != nil {
			fmt.Println(err)
		}
		psc := redis.PubSubConn{c}

		// Set up subscriptions
		psc.Subscribe("__keyevent@0__:expired")

		// While not a permanent error on the connection.
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
}
