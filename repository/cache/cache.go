package cache

import (
	"context"
	"iserver/obj"
	"iserver/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetLastLoginAccountJwtToken(account, jwtToken string) error {
	if err := RedisClient.Set(context.Background(),
		bindLastLoginTokenKey(account), jwtToken, time.Hour*24).Err(); err != nil {
		return err
	}

	return nil
}

func GetLastLoginAccountJwtToken(account string) (string, error) {
	token, err := RedisClient.Get(context.Background(), bindLastLoginTokenKey(account)).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(account string) (obj.UserData, error) {
	result, err := RedisClient.Get(context.Background(), bindUserInfoKey(account)).Bytes()
	if err != nil {
		return obj.UserData{}, err
	}

	data := obj.UserData{}

	utils.Json.Unmarshal(result, &data)

	return data, nil
}

func SetUser(account string, data obj.UserData) error {
	b, _ := utils.Json.Marshal(data)

	if err := RedisClient.Set(context.Background(), bindUserInfoKey(account), string(b), time.Hour).Err(); err != nil {
		return err
	}

	return nil
}
