package main

import (
	"github.com/itziklavon/kit2go/common/src/common"
	"fmt"
)

func main() {
	fmt.Printf(common.GetRedisConnection(7))
	fmt.Println()
}
