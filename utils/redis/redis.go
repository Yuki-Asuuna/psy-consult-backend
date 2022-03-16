package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"psy-consult-backend/database"
)

const (
	redis_address  = "127.0.0.1:6379"
	redis_password = ""
	redis_network  = "tcp"
	expire_time    = 3600
)

var client *redis.Client
var ctx = context.Background()

func RedisInit() {
	client = redis.NewClient(&redis.Options{
		Addr:     redis_address,
		Password: redis_password,
		Network:  redis_network,
		DB:       1, // 仓库编号
	})
}

func GetRedisClient() *redis.Client {
	return client
}

func SetWxSessionKey(key string, openID string) error {
	err := client.Set(ctx, key, openID, expire_time).Err()
	if err != nil {
		logrus.Errorf("SetWxSessionKey failed, err= %v", err)
		return err
	}
	return nil
}

func GetWxOpenIDBySessionKey(sessionKey string) (string, error) {
	val, err := client.Get(ctx, sessionKey).Result()
	if err != nil {
		logrus.Errorf("GetWxOpenIDBySessionKey failed, err= %v", err)
		return "", err
	}
	return val, nil
}

func GetVisitorInfoBySessionKey(sessionKey string) *database.VisitorUser {
	openID, err := GetWxOpenIDBySessionKey(sessionKey)
	if err != nil {
		return nil
	}
	user, err := database.GetVisitorUserByVisitorID(openID)
	if err != nil {
		return nil
	}
	if user == nil {
		return nil
	}
	return user
}
