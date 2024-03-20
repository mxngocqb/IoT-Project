package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/mxngocqb/IoT-Project/config"
	"github.com/mxngocqb/IoT-Project/model"
	"github.com/redis/go-redis/v9"
)

var ctx context.Context = context.Background()

type vehicleCacheImpl struct {
	client  *redis.Client
	expires time.Duration
}

func (vehicleCache *vehicleCacheImpl) Delete(key string) error {
	redisKey := "vehicle:" + key
	return vehicleCache.client.Del(ctx, redisKey).Err()
}

func NewVehicleRedisCache(client *redis.Client) VehicleCache {
	expires, _ := time.ParseDuration(config.VehicleCacheExpirationMs)

	return &vehicleCacheImpl{
		client:  client,
		expires: expires,
	}
}

func (vehicleCache *vehicleCacheImpl) Set(key string, value *model.Vehicle) {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	redisKey := "vehicle:" + key
	vehicleCache.client.Set(ctx, redisKey, json, vehicleCache.expires)
}

func (vehicleCache *vehicleCacheImpl) Get(key string) *model.Vehicle {

	redisKey := "vehicle:" + key
	val, err := vehicleCache.client.Get(ctx, redisKey).Result()
	if err != nil {
		return nil
	}

	vehicle := model.Vehicle{}
	err = json.Unmarshal([]byte(val), &vehicle)
	if err != nil {
		panic(err)
	}
	return &vehicle
}

func (vehicleCache *vehicleCacheImpl) GetMultiRequest(key string) []model.Vehicle {
	// Retrieve data from Redis
	data, err := vehicleCache.client.Get(ctx, key).Result()
	if err != nil {
		log.Println("Error getting data from Redis:", err)
		return nil
	}

	// Decode JSON
	var vehicles []model.Vehicle
	if err := json.Unmarshal([]byte(data), &vehicles); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return vehicles
}

func (vehicleCache *vehicleCacheImpl) SetMultiRequest(key string, value []model.Vehicle)  {
	// Encode slice of vehicle objects to JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		panic(err)
	}

	// Save to Redis
	err = vehicleCache.client.Set(ctx, key, jsonData, vehicleCache.expires).Err()
	if err != nil {
		log.Println("Error setting data in Redis:", err)
		panic(err)
	}
}
