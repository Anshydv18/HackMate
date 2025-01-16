package redisentity

import (
	base "Hackmate/Base"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const EmailOtpKey = "hackmate::email"

func SetOtpCache(ctx *context.Context, email string, otp int) error {
	MainKey := fmt.Sprintf("%s:%s", EmailOtpKey, email)
	rdb := base.RedisInstance
	Otp, err := json.Marshal(otp)
	if err != nil {
		return err
	}
	res := rdb.Set(*ctx, MainKey, Otp, 5*time.Minute)
	fmt.Println(res)
	return nil
}

func GetOtpCache(ctx *context.Context, email string) (int, error) {
	MainKey := fmt.Sprintf("%s:%s", EmailOtpKey, email)
	rdb := base.RedisInstance
	value := rdb.Get(*ctx, MainKey)
	otp, _ := value.Int()
	return otp, nil
}
