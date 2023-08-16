package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func RdbConnect() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6378",
		Password: "123456",
		DB:       0,
	})
	ctx := context.Background()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err.Error())
	}

	//Rdb.Set(ctx, "name", "zhou", 600)
	//res := Rdb.Get(ctx, "name")
	//fmt.Println(res.Val())
}
