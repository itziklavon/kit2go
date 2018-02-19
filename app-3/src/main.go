package main

import "github.com/itziklavon/kit2go/redis-helper/src/redis_helper"

func main() {
	fmt.Printf(redis.GetRedisConnection(7))
	fmt.Println()
}
