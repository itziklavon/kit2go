package mysql_helper

import (
	"database/sql"

	"github.com/itziklavon/kit2go/general-log/src/general_log"

	"github.com/itziklavon/kit2go/configuration/src/configuration"

	_ "github.com/go-sql-driver/mysql"
	"github.com/itziklavon/kit2go/http-client-helper/src/http_client_helper"
	"github.com/jmoiron/sqlx"
)

func GetMultiBrandConnection(discoveryData http_client_helper.DiscoveryDbData) *sqlx.DB {
	mysqluri := discoveryData.UserName + ":" +
		discoveryData.Password + "@tcp(" + discoveryData.Host + ":3306" +
		")/" + discoveryData.SchemaName
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/story")
	general_log.Debug(mysqluri)
	if err != nil {
		general_log.ErrorException(":GetMultiBrandConnection: couldn't connect to DB, host: "+discoveryData.Host, err)
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
		general_log.ErrorException(":GetMultiBrandConnection: couldn't connect to local db", err)
	}
	return db
}
