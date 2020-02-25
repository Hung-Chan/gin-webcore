package redis

import (
	"fmt"

	"os"

	"github.com/go-redis/redis/v7"
)

// RedisManage .
var RedisManage *redis.Client

func init() {
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")
	port := os.Getenv("REDIS_PORT")

	RedisManage = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	pong, err := RedisManage.Ping().Result()

	fmt.Println(pong, err)
}

// SetValue .
func SetValue(key, value string, second int64) error {

	err := RedisManage.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// DeleteValue .
func DeleteValue(key string) error {

	err := RedisManage.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
