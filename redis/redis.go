package main

import (
	"fmt"
	// "reflect"
	"github.com/gomodule/redigo/redis"
)

func main() {

	// 打開redis
	redisConnect, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("somthing error")
	}
	defer redisConnect.Close()

	// fmt.Println(reflect.TypeOf(redisConnect))

	// 寫入redis
	setToRedis(redisConnect, "123")

	// 讀取redis
	value := getFromRedis(redisConnect)

	fmt.Println(value)

}

// 寫入redis
func setToRedis(redisC redis.Conn, content string) {
	// fmt.Println(content)
	redisC.Do("SET", "aduit_key", content)

}

// 讀取redis
func getFromRedis(redisC redis.Conn) string {
	value, _ := redis.String(redisC.Do("GET", "aduit_key"))
	return value
}
