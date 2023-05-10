package api

// var client

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/pouyam79i/Cloud_Computing_HW/main/HW2/step2/code/util"
)

var client *redis.Client = nil

func initialConnectionToRedis() error {
	server_info, err := util.GetConfigs()
	if err != nil {
		return err
	}
	client = redis.NewClient(&redis.Options{
		Addr:     server_info.REDIS_ADDR,
		Password: "",
		DB:       0,
	})
	return nil
}

// It returns shortened URL or error
func GetRedis(key string) (string, error) {
	if client == nil {
		err := initialConnectionToRedis()
		if err != nil {
			return "", err
		}
	}
	if client == nil {
		return "", errors.New("nil redis client")
	}
	if err := client.Ping(); err != nil {
		return "", errors.New("no redis client connection")
	}
	data := client.Get(key)
	if errors.Is(data.Err(), redis.Nil) {
		return "", redis.Nil
	}
	return data.Val(), nil
}

// It save shortened URL on redis
func SendRedis(key, val string) error {
	if client == nil {
		err := initialConnectionToRedis()
		if err != nil {
			return err
		}
	}
	if client == nil {
		return errors.New("nil redis client")
	}
	if err := client.Ping(); err != nil {
		return errors.New("no redis client connection")
	}
	exp := time.Duration(300 * time.Second) // 5 minutes
	cmd := client.Set(key, val, exp)
	return cmd.Err()
}