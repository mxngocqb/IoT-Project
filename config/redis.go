package config

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

// RedisConfiginterface redis config interface
type RedisConfiginterface interface {
	Address() string
	Password() string
}

// Redis redis struct
type Redis struct {
	address  string `env:"REDIS_ADDRESS" envDefault:"localhost:6379"`
	password string `env:"REDIS_PASSWORD" envDefault:""`
}

// NewRedisConfig create redis instance
func NewRedisConfig() *Redis {
	address := "localhost:6379"
	password := ""

	if env := os.Getenv("REDIS_ADDRESS"); env != "" {
		address = env
	}
	if env := os.Getenv("REDIS_PASSWORD"); env != "" {
		password = env
	}

	redis := &Redis{
		address:  address,
		password: password,
	}
	return redis
}

// Address get redis address
func (redis *Redis) Address() string {
	return redis.address
}

// Password get redis password
func (redis *Redis) Password() string {
	return redis.password
}

// ConnectRedis connects to Redis using provided configuration
func ConnectRedis(config RedisConfiginterface) (*redis.Client, error) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address(),
		Password: config.Password(),
		DB:       0, // default DB
	})

	res, err := rdb.Ping(ctx).Result()

	if err != nil {
		fmt.Println("There is an error while connecting to the Redis ", err)
		return nil, err
	} else {
		fmt.Println("Redis connected", res)
	}

	return rdb, nil
}
