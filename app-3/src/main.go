package main

import (
	"github.com/itziklavon/kit2go/redis-helper/src/redis_helper"
)

func main() {
	fmt.Printf(redis_helper.GetRedisConnection(7))
	fmt.Println()
}
