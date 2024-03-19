package config

import "os"

var (
	DriverCacheExpirationMs = GetEnv("Driver_CACHE_EXPIRATION_MS", "5000ms")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
