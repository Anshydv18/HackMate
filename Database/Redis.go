package database

import (
	constants "NotesBuddy/Constants"
	env "NotesBuddy/Env"

	"github.com/redis/go-redis/v9"
)

func StartRedisServer() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.Get(constants.RD_CONN),
		Username: "default",
		Password: env.Get(constants.RD_PASS),
		DB:       0,
	})

	return rdb
}