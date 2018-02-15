package common

import (
	"general_log"
)

func PrintHello(name string) {
	general_log.SetLogOutput("/hello/hello.log")
	general_log.Debug("Hello," + name)
}
