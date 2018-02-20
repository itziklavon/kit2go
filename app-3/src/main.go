package main

import (
	"github.com/itziklavon/kit2go/redis-common/src"
)

func main() {
	fmt.Printf(redis_common.GetRedisConnection(7))
	fmt.Println()
}
