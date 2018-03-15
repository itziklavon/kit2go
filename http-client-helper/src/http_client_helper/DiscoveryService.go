package http_client_helper

import (
	"encoding/json"

	"github.com/itziklavon/kit2go/configuration/src/configuration"
)

var DISCOVERY_URL = configuration.GetPropertyValue("DISCOVERY_URL")

var discoveryMapping = map[int]map[string]string{}

type DiscoveryData struct {
	BrandId     int    `json:"brand_id"`
	ServiceName string `json:"service_name"`
	Url         string `json:"url"`
}

type ServiceData struct {
	ServiceName string `json:"service_name"`
	Url         string `json:"url"`
}

func GetDiscoveryUrl(brandId int, serviceName string) string {
	if len(discoveryMapping) == 0 || (len(discoveryMapping[-1][serviceName]) == 0 && len(discoveryMapping[brandId][serviceName]) == 0) {
		initMap()
	}
	if len(discoveryMapping[brandId][serviceName]) == 0 {
		return discoveryMapping[-1][serviceName]
	}
	return discoveryMapping[brandId][serviceName]
}

func initMap() {
	discoveryMapping = make(map[int]map[string]string)
	url := DISCOVERY_URL + "discovery-web/services"
	httpResponse := GETBody(url, nil)
	var arr []DiscoveryData
	_ = json.Unmarshal([]byte(httpResponse), &arr)
	for i := 0; i < len(arr); i = i + 1 {
		if discoveryMapping[arr[i].BrandId] == nil {
			discoveryMapping[arr[i].BrandId] = make(map[string]string)
		}
		serviceData := ServiceData{ServiceName: arr[i].ServiceName, Url: arr[i].Url}
		discoveryMapping[arr[i].BrandId][serviceData.ServiceName] = serviceData.Url
	}
}
