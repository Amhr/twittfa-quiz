package rdb

import (
	"WhoKnowsMeapp/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.RDB_HOST, config.RDB_PORT),
		DB:   0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
