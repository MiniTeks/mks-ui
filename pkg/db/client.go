package db

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func GetRedisClient(cred *RClient) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cred.Addr,
		Password: cred.Pass, // no password set
		DB:       cred.Db,   // use default DB 0
	})

	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Couldn't connect to redis-database")
	}
	fmt.Print(ping, err)

	return rdb
}
