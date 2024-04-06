package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisCache) Set(key string, value string) error {
	ctx := context.Background()
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisCache) SetWithExpire(key string, value string, ttl uint64) error {
	ctx := context.Background()
	return r.client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

func (r *RedisCache) Delete(key string) error {
	ctx := context.Background()
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) Flush() error {
	ctx := context.Background()
	return r.client.FlushAll(ctx).Err()
}
