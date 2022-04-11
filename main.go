package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type redis struct {
	Database map[string]string `json:"database"`
	sync.RWMutex
}

var list = redis{Database: map[string]string{
	"1": "armanchik",
	"2": "Hello, world!",
	"3": "Danial hello bratiwka",
}}

func getList(context *gin.Context) {
	list.RLock()
	key := context.Param("key")
	list.RUnlock()
	id, val := getListById(key)

	if val == "err" {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": id})
		return
	}
	context.IndentedJSON(http.StatusOK, "Key : "+id+", Value : "+val)
}

func getListById(key string) (string, string) {
	for i, t := range list.Database {
		if i == key {
			return i, t
		}
	}

	return "Key not found", "err"
}

func updateList(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")

	id, val := updateValueInList(key, value)

	context.IndentedJSON(http.StatusOK, "Key : "+id+", Value : "+val)
}

func updateValueInList(key string, value string) (string, string) {
	list.Lock()
	list.Database[key] = value
	list.Unlock()
	return key, value
}

func main() {
	router := gin.Default()
	router.GET("/store/:key", getList)
	router.PUT("/store/:key/:value", updateList)
	router.Run("localhost:9090")
}
