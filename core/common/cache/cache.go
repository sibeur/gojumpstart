package cache

import (
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheUseCase interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	SetWithExpire(key string, value string, ttl uint64) error
	Delete(key string) error
	Flush() error
}

func NewCache() CacheUseCase {
	var err error
	var cache CacheUseCase
	// discover cache type
	cacheType := os.Getenv("CACHE_TYPE")
	if cacheType == "" {
		cacheType = "redis"
	}

	switch cacheType {
	case "redis":
		// load redis
		redisDb := 0
		if os.Getenv("REDIS_DB") != "" {
			redisDb, err = strconv.Atoi(os.Getenv("REDIS_DB"))
			if err != nil {
				panic(err)
			}
		}
		redisClient := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"), // no password set
			DB:       redisDb,                     // use default DB
		})

		// load cache
		cache = NewRedisCache(redisClient)

	case "memory":
		// load memory cache
		cache = NewMemoryCache(time.Hour * 1)
	default:
		panic("Cache type not supported")
	}
	return cache
}
