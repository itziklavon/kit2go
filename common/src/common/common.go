package common

import (
	"github.com/itziklavon/kit2go/general-log/src/general_log"
	"github.com/itziklavon/kit2go/configuration/src/configuration"
	"github.com/go-redis/redis"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"net/http"
)

var redisHost = configuration.GetPropertyValue("REDIS_HOST")

var DISCOVERY_URL = configuration.GetPropertyValue("DISCOVERY_URL")

var redisConnections map[int]string

type RedisData struct {
	BrandId int `json:"brandId"`
	Uri     string `json:"uri"`
}

type GenericHttpResponse struct {
	httpResponse string
	httpBody     string
	httpHeaders  map[string]string
}

func PrintHello(name string) {
	general_log.SetLogOutput("/hello/hello.log")
	general_log.Debug("Hello," + name)
}

func GET(url string, headers map[string]string) GenericHttpResponse {
	req, err := http.NewRequest("GET", url, nil)
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		general_log.ErrorException(":GET: couldn't send request", err)
	}
	defer resp.Body.Close()
	general_log.Info("response Status:" + resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	general_log.Info("response Body:" + string(body))
	var responseHeaders = make(map[string]string)
	for key, value := range resp.Header {
		responseHeaders[key] = value[0]
		general_log.Info(":GET: response Headers: key: " + key + ", value: " + value[0])
	}
	return GenericHttpResponse{httpResponse: resp.Status, httpBody: string(body), httpHeaders: responseHeaders}
}

func GETBody(url string, headers map[string]string) string {
	req, err := http.NewRequest("GET", url, nil)
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		general_log.ErrorException(":GET: couldn't send request", err)
	}
	defer resp.Body.Close()
	general_log.Info("response Status:" + resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	general_log.Info("response Body:" + string(body))
	var responseHeaders = make(map[string]string)
	for key, value := range resp.Header {
		responseHeaders[key] = value[0]
		general_log.Info(":GET: response Headers: key: " + key + ", value: " + value[0])
	}
	return string(body)
}

func POST(url string, body map[string]string, headers map[string]string) GenericHttpResponse {
	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		general_log.ErrorException(":POST: couldn't send request", err)
	}
	defer resp.Body.Close()
	general_log.Info("response Status:" + resp.Status)
	responseBody, _ := ioutil.ReadAll(resp.Body)
	general_log.Info("response Body:" + string(responseBody))
	var responseHeaders = make(map[string]string)
	for key, value := range resp.Header {
		responseHeaders[key] = value[0]
		general_log.Info(":POST: response Headers: key: " + key + ", value: " + value[0])
	}
	return GenericHttpResponse{httpResponse: resp.Status, httpBody: string(responseBody), httpHeaders: responseHeaders}
}

func PUT(url string, body map[string]string, headers map[string]string) GenericHttpResponse {
	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		general_log.ErrorException(":PUT: couldn't send request", err)
	}
	defer resp.Body.Close()
	general_log.Info("response Status:" + resp.Status)
	responseBody, _ := ioutil.ReadAll(resp.Body)
	general_log.Info("response Body:" + string(responseBody))
	var responseHeaders = make(map[string]string)
	for key, value := range resp.Header {
		responseHeaders[key] = value[0]
		general_log.Info(":PUT: response Headers: key: " + key + ", value: " + value[0])
	}
	return GenericHttpResponse{httpResponse: resp.Status, httpBody: string(responseBody), httpHeaders: responseHeaders}
}

func DELETE(url string, body map[string]string, headers map[string]string) GenericHttpResponse {
	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonValue))
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		general_log.ErrorException(":DELETEGET: couldn't send request", err)
	}
	defer resp.Body.Close()
	general_log.Info("response Status:" + resp.Status)
	responseBody, _ := ioutil.ReadAll(resp.Body)
	general_log.Info("response Body:" + string(responseBody))
	var responseHeaders = make(map[string]string)
	for key, value := range resp.Header {
		responseHeaders[key] = value[0]
		general_log.Info(":DELETE: response Headers: key: " + key + ", value: " + value[0])
	}
	return GenericHttpResponse{httpResponse: resp.Status, httpBody: string(responseBody), httpHeaders: responseHeaders}
}

func GetRedisConnection(brandId int) string {
	if len(redisConnections) == 0 {
		initMap()
	}
	return redisConnections[brandId]
}

func initMap() {
	redisConnections = make(map[int]string)
	url := DISCOVERY_URL + "discovery-web/brand/services/REDIS"
	httpResponse := GETBody(url, nil)
	var arr []RedisData
	_ = json.Unmarshal([]byte(httpResponse), &arr)
	for i := 0; i < len(arr); i = i + 1 {
		redisConnections[arr[i].BrandId] = arr[i].Uri
	}
}

func getRedisConnection() *redis.Client {
	return getRedisConnectionByHost(redisHost)
}

func getRedisConnectionByHost(host string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		general_log.ErrorException(":getRedisConnection: couldn't connect ro redis", err)
	}
	return client
}

func Keys() []string {
	conn := getRedisConnection()
	value, err := conn.Keys("*").Result()
	if err != nil {
		general_log.ErrorException(":Keys: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func KeysWithPattern(pattern string) []string {
	conn := getRedisConnection()
	value, err := conn.Keys(pattern).Result()
	if err != nil {
		general_log.ErrorException(":Keys: couldn't get Keys from redis", err)
	}
	defer conn.Close()
	return value
}

func Get(key string) string {
	conn := getRedisConnection()
	value, err := conn.Get(key).Result()
	if err != nil {
		general_log.ErrorException(":Get: couldn't get key from redis: " + key, err)
	}
	defer conn.Close()
	return value
}

//func Set(key string, value string) string {
//	conn := getRedisConnection()
//	str, err := conn.Set(key, value).Result()
//	if err != nil {
//		general_log.ErrorException(":Set: couldn't set key from redis: " + key, err)
//	}
//	defer conn.Close()
//	return str
//}

func HGet(key string, hkey string) string {
	conn := getRedisConnection()
	value, err := conn.HGet(key, hkey).Result()
	if err != nil {
		general_log.ErrorException(":Set: couldn't get key from redis: " + key + ", hKey: " + hkey, err)
	}
	defer conn.Close()
	return value
}

//func HSet(key string, hkey string, value string) {
//	conn := getRedisConnection()
//	str, err := conn.HSet(key, hkey, value).Result()
//	if str {
//		general_log.ErrorException(":Set: key doesn't exists in redis: " + key + ", hKey: " + hkey, err)
//	}
//	if err != nil {
//		general_log.ErrorException(":Set: couldn't get key from redis: " + key + ", hKey: " + hkey, err)
//	}
//	defer conn.Close()
//}

func GetSysParam(hkey string) string {
	return HGet("SysParams", hkey)
}

func GetBrandId() string {
	return GetSysParam("GS_BRAND_ID")
}

