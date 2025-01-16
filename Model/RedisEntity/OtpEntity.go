package redisentity

import (
	base "Hackmate/Base"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const EmailOtpKey = "hackmate::email"

func SetOtpCache(ctx *context.Context, email string, otp int) {
	MainKey := fmt.Sprintf("%s:%s", EmailOtpKey, email)
	rdb := base.RedisInstance
	Otp, _ := json.Marshal(otp)
	rdb.Set(*ctx, MainKey, Otp, 5*time.Second)
}

func GetOtpCache(ctx *context.Context, email string) (int, error) {
	MainKey := fmt.Sprintf("%s:%s", EmailOtpKey, email)
	rdb := base.RedisInstance
	value := rdb.Get(*ctx, MainKey)

	fmt.Println(value)

	var otp int
	if err := json.Unmarshal([]byte(value.String()), &otp); err != nil {
		return 0, err
	}
	return otp, nil
}
