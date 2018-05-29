package main

import (
	"github.com/itziklavon/kit2go/configuration/src/configuration"
	"log"
	//"github.com/itziklavon/kit2go/redis/src/redis"
)

func main() {
	logoutApi := configuration.GetTogglesPropertyValue("LOGOUT_API")
	log.Println(logoutApi)
}
