package http_client_helper

import (
	"encoding/base64"
	"encoding/json"

	"github.com/itziklavon/kit2go/configuration/src/configuration"
)

var discoveryDbMapping map[int]DiscoveryDbData

type DiscoveryDbData struct {
	BrandId    int    `json:"brandId"`
	SchemaName string `json:"schemaName"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
}

func GetDiscoveryDbConnection(brandId int) DiscoveryDbData {
	if len(discoveryDbMapping) == 0 {
		initDbMap()
	}
	if discoveryDbMapping[brandId] == (DiscoveryDbData{}) {
		initDbMap()
	}
	return discoveryDbMapping[brandId]
}

func initDbMap() {
	discoveryDbMapping = make(map[int]DiscoveryDbData)
	url := DISCOVERY_URL + "discovery-web/brand/database/connector"
	httpResponse := GETBody(url, nil)
	var arr []DiscoveryDbData
	_ = json.Unmarshal([]byte(httpResponse), &arr)
	for i := 0; i < len(arr); i = i + 1 {
		serviceData := DiscoveryDbData{BrandId: arr[i].BrandId, SchemaName: arr[i].SchemaName,
			UserName: arr[i].UserName, Password: decodePassword(arr[i].Password), Host: arr[i].Host}
		discoveryDbMapping[arr[i].BrandId] = serviceData
	}
}

func decodePassword(encryptedPass string) string {
	str, _ := base64.StdEncoding.DecodeString(encryptedPass)
	decryptedPass := DecryptPassword(str, configuration.GetPropertyValue("ENCRYPTION_PRIVATE_KEY"))
	return string(decryptedPass[:len(decryptedPass)-2])
}
