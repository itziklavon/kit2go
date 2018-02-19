package main

import (
	"github.com/itziklavon/kit2go/common/src/common"
	"github.com/itziklavon/kit2go/redis/src/redis"
)

func main() {
	common.PrintHello("world")
	common.PrintHello(redis.GGetRedisConnection(7))
}
