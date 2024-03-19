package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mxngocqb/IoT-Project/config"
	"github.com/mxngocqb/IoT-Project/model"
	"github.com/redis/go-redis/v9"
)

var ctx context.Context = context.Background()

type driverCacheImpl struct {
	client  *redis.Client
	expires time.Duration
}

func (driverCache *driverCacheImpl) Delete(key string) error {
	return driverCache.client.Del(ctx, key).Err()
}

func NewDriverRedisCache(client *redis.Client) DriverCache {
	expires, _ := time.ParseDuration(config.DriverCacheExpirationMs)

	return &driverCacheImpl{
		client:  client,
		expires: expires,
	}
}

func (driverCache *driverCacheImpl) Set(key string, value *model.Driver) {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	driverCache.client.Set(ctx, key, json, driverCache.expires)
}

func (driverCache *driverCacheImpl) Get(key string) *model.Driver {

	val, err := driverCache.client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	driver := model.Driver{}
	err = json.Unmarshal([]byte(val), &driver)
	if err != nil {
		panic(err)
	}
	return &driver
}
