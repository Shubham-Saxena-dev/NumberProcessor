package cache

import (
	"CARIAD/config"
	"CARIAD/internal/customerrors"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"sync"
	"time"
)

var (
	once  sync.Once
	cache RedisCache
)

const DefaultExpiration time.Duration = 0

type RedisCache interface {
	Get(key interface{}) (string, error)
	Set(key interface{}, value string) error
	Delete(key interface{}) error
	DeleteAll() error
}

type redisClient struct {
	client *redis.Client
	json   jsoniter.API
}

func GetRedisCache() RedisCache {

	once.Do(func() {
		redisConfig := config.EnvConfigs.Redis
		redClient := &redisClient{
			client: redis.NewClient(
				&redis.Options{
					Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
					Password: redisConfig.Password,
					DB:       redisConfig.DB,
				}),
			json: jsoniter.ConfigCompatibleWithStandardLibrary,
		}
		cache = redClient
	})

	return cache
}

func (rc *redisClient) Get(key interface{}) (string, error) {
	val, err := rc.client.Get(context.Background(), rc.getRedisKey(key)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", customerrors.ErrorInvalidKey(key, err)
		}
		return "", customerrors.ErrorRedisGet(key, err)
	}
	return val, nil
}

func (rc *redisClient) Set(key interface{}, value string) error {
	err := rc.client.Set(context.Background(), rc.getRedisKey(key), value, DefaultExpiration).Err()
	if err != nil {
		return customerrors.ErrorRedisSet(key, err)
	}
	return nil
}

func (rc *redisClient) Delete(key interface{}) error {
	err := rc.client.Del(context.Background(), rc.getRedisKey(key)).Err()
	if err != nil {
		return customerrors.ErrorRedisDel(key, err)
	}
	return err
}

func (rc *redisClient) DeleteAll() error {
	_, err := rc.client.FlushDB(context.Background()).Result()
	if err != nil {
		return customerrors.ErrorDelAll(err)
	}
	return err
}

func (rc *redisClient) getRedisKey(key interface{}) string {
	if strKey, ok := key.(string); ok {
		return strKey
	}
	jsonBytes, _ := rc.json.Marshal(key)
	return string(jsonBytes)
}
