package mysql_helper

import (
	"database/sql"
	"log"

	"github.com/itziklavon/kit2go/configuration/src/configuration"

	_ "github.com/go-sql-driver/mysql"
	"github.com/itziklavon/kit2go/http-client-helper/src/http_client_helper"
)

func GetMultiBrandConnection(discoveryData http_client_helper.DiscoveryDbData) *sql.DB {
	mysqluri := discoveryData.UserName + ":" +
		discoveryData.Password + "@tcp(" + discoveryData.Host + ":3306" +
		")/" + discoveryData.SchemaName
	log.Println(mysqluri)
	db, err := sql.Open("mysql", mysqluri)
	if err != nil {
		log.Println(":GetMultiBrandConnection: couldn't connect to DB, host: " + discoveryData.Host)
		log.Println(err)
	}
	return db
}

func GetLocalConnection(schemaName string) *sql.DB {
	username := configuration.GetPropertyValue("JDBC_USER_NAME")
	password := configuration.GetPropertyValue("JDBC_PASSWORD")
	host := configuration.GetPropertyValue("JDBC_HOST")
	db, err := sql.Open("mysql", username+":"+
		password+"@tcp("+host+
		")/"+schemaName)
	if err != nil {
		log.Println(":GetLocalConnection: couldn't connect to DB, host: " + host)
		log.Println(err)
	}
	return db
}
