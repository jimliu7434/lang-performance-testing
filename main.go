package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

// PostDta is used to get post data
type PostDta struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	fmt.Printf("Hello world\r\n")
	createClient()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/redis/:key", getRedis)
	router.POST("/redis", postRedis)

	s := &http.Server{
		Addr:           ":7001",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func getRedis(c *gin.Context) {
	key := c.Param("key")
	result := getRedisValue(key)

	c.JSON(http.StatusOK, gin.H{
		"key":   key,
		"value": result,
	})
}

func postRedis(c *gin.Context) {
	var postData PostDta
	c.BindJSON(&postData)
	key := postData.Key
	value := postData.Value
	setRedisValue(key, value)
	result := "Set to redis success with key: " + key + ", and with value: " + getRedisValue(key)
	c.String(http.StatusOK, result)
}

func createClient() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "172.16.88.145:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func setRedisValue(key, value string) {
	err := redisdb.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func getRedisValue(key string) string {
	val, err := redisdb.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
	return val
}
