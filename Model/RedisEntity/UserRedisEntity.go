package redisentity

import (
	database "NotesBuddy/Database"
	dto "NotesBuddy/Model/Dto"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const UserRedisKey = "user"

func SetUserCache(ctx *context.Context, data *dto.User) {
	key := fmt.Sprintf("%s:%s", UserRedisKey, data.Phone)
	rdb := database.StartRedisServer()
	jsonUser, _ := json.Marshal(data)
	rdb.Set(*ctx, key, jsonUser, 30*time.Minute)
}

func GetUserFromCache(ctx *context.Context, phone string) (*dto.User, error) {
	key := fmt.Sprintf("%s:%s", UserRedisKey, phone)
	rdb := database.StartRedisServer()
	result := rdb.Get(*ctx, key)

	var UserDto *dto.User
	err := json.Unmarshal([]byte(result.String()), &UserDto)
	return UserDto, err
}
