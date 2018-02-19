package main

import "github.com/itziklavon/kit2go/redis/src/redis"

func main() {
	fmt.Printf(redis.GetRedisConnection(7))
	fmt.Println()
}
