package database

import "github.com/go-redis/redis/v8"



func NewRedis()*redis.Client{ 
	rdb := redis.NewClient(&redis.Options{
        Addr:     "cache:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	return rdb
}