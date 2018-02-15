package common

import (
	"github.com/itziklavon/kit2go/general-log/src/general_log"
)

func PrintHello(name string) {
	general_log.SetLogOutput("/hello/hello.log")
	general_log.Debug("Hello," + name)
}
