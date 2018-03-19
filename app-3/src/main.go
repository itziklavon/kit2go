package main

import (
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/itziklavon/kit2go/configuration/src/configuration"
	"github.com/itziklavon/kit2go/general-log/src/general_log"
	"github.com/itziklavon/kit2go/http-client-helper/src/http_client_helper"
	"github.com/itziklavon/kit2go/mysql-helper/src/mysql_helper"
	"github.com/itziklavon/kit2go/redis-helper/src/redis_helper"
)

type Row struct {
	ID int `json:"id"`
}

var fileName = configuration.GetLogPropertyValue("REDIS_LISTENER_LOG")

var done = make(chan bool)

func main() {
	general_log.SetLogOutput(fileName)
	brandSet := redis_helper.GetBrandSet()
	general_log.Debug(":main: extracted brands are: ", brandSet)
	for _, brandId := range brandSet {
		general_log.Debug(":main: staring for brand: ", brandId)
		go handleMessges(brandId)
	}
	general_log.Debug(strconv.FormatBool(<-done))
}

func handleMessges(brandId int) {
	for {
		redisHelper := redis_helper.GetRedisConnection(brandId)
		c, psc := redisHelper.Subscribe("__keyevent@0__:expired")
		for c.Err() == nil {
			switch v := psc.Receive().(type) {
			case redis.Message:
				messageData := string(v.Data[:])
				general_log.Debug(":handleMessges: exracted message is: " + messageData)
				if strings.HasPrefix(messageData, "TKN_") {
					if redisHelper.Exists(messageData[4:len(messageData)]) {
						playerId := redisHelper.HGet(messageData[4:len(messageData)], "id")
						go logoutPlayer(brandId, messageData[4:len(messageData)], playerId)
					}
				}
			}
		}
		general_log.ErrorException(":handleMessges: an error occured in brand: "+string(brandId), c.Err())
		c.Close()
	}
	done <- true
}

func logoutPlayer(brandId int, token string, playerId string) {
	stdBrand := strconv.Itoa(brandId)
	url := http_client_helper.GetDiscoveryUrl(brandId, "LOGOUT") + "/1.29.08/" + stdBrand + "/player"
	values := map[string]string{"auth_token": token}
	headers := map[string]string{"x-auth-token": token}
	general_log.Debug("sending message to brand ", brandId, " with token:"+token, "to uri:"+url)
	http_client_helper.POST(url, values, headers)
	db := mysql_helper.GetMultiBrandConnection(http_client_helper.GetDiscoveryDbConnection(7))
	defer db.Close()
	general_log.Debug("isnerting audit log for player: " + playerId)

	auditLogInsert := "INSERT INTO tbl_auditLog (`playerId`, `date`, `action`, `operator`, `comment`, `newData`, `oldData`, `fieldName`) VALUES (?, now(), 'player_session_expired_logout', 'player', 'logging out player - session expired, automatic logout', 'logged-out', 'logged-in', 'Login status')"
	stmtIns, err := db.Prepare(auditLogInsert)
	if err != nil {
		general_log.ErrorException("an error occurred", err)
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(playerId)
	if err != nil {
		general_log.ErrorException("an error occurred", err)
	}
}
