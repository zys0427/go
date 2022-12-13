package service

import (
	"fmt"
	"github.com/go-redis/redis"
)

type RedisZys struct {

}

// RedisConnZys redis 连接
func (RedisZys RedisZys) RedisConnZys() redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return *client
}
