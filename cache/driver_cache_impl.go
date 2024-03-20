package cache

import (
	// "context"
	"encoding/json"
	"log"
	"time"

	"github.com/mxngocqb/IoT-Project/config"
	"github.com/mxngocqb/IoT-Project/model"
	"github.com/redis/go-redis/v9"
)

// var ctx context.Context = context.Background()

type driverCacheImpl struct {
	client  *redis.Client
	expires time.Duration
}

func (driverCache *driverCacheImpl) Delete(key string) error {
	redisKey := "driver:" + key
	return driverCache.client.Del(ctx, redisKey).Err()
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
	redisKey := "driver:" + key
	driverCache.client.Set(ctx, redisKey, json, driverCache.expires)
}

func (driverCache *driverCacheImpl) Get(key string) *model.Driver {

	redisKey := "driver:" + key
	val, err := driverCache.client.Get(ctx, redisKey).Result()
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

func (driverCache *driverCacheImpl) GetMultiRequest(key string) []model.Driver {
	// Retrieve data from Redis
	data, err := driverCache.client.Get(ctx, key).Result()
	if err != nil {
		log.Println("Error getting data from Redis:", err)
		return nil
	}

	// Decode JSON
	var drivers []model.Driver
	if err := json.Unmarshal([]byte(data), &drivers); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return drivers
}

func (driverCache *driverCacheImpl) SetMultiRequest(key string, value []model.Driver)  {
	// Encode slice of Driver objects to JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		panic(err)
	}

	// Save to Redis
	err = driverCache.client.Set(ctx, key, jsonData, driverCache.expires).Err()
	if err != nil {
		log.Println("Error setting data in Redis:", err)
		panic(err)
	}
}
