package redisentity

import (
	base "Hackmate/Base"
	dto "Hackmate/Model/Dto"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const UserRedisKey = "hackmate::user"

func SetUserCache(ctx *context.Context, Key string, data *dto.User) {
	key := fmt.Sprintf("%s:%s", UserRedisKey, Key)
	rdb := base.RedisInstance
	jsonUser, _ := json.Marshal(data)
	rdb.Set(*ctx, key, jsonUser, 30*time.Minute)
}

func GetUserFromCache(ctx *context.Context, Key string) (*dto.User, error) {
	key := fmt.Sprintf("%s:%s", UserRedisKey, Key)
	rdb := base.RedisInstance
	result := rdb.Get(*ctx, key)

	var UserDto *dto.User
	err := json.Unmarshal([]byte(result.String()), &UserDto)
	return UserDto, err
}
