package cache_test

import (
	core_cache "gojumpstart/core/common/cache"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRedisCache_Get(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := core_cache.NewRedisCache(client)

	// Set a value in the cache
	err := cache.Set("key", "value")
	if err != nil {
		t.Errorf("Failed to set value in cache: %v", err)
	}

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value matches the expected value
	expectedValue := "value"
	if value != expectedValue {
		t.Errorf("Expected value %s, but got %s", expectedValue, value)
	}
}

func TestRedisCache_Set(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := core_cache.NewRedisCache(client)

	// Set a value in the cache
	err := cache.Set("key", "value")
	if err != nil {
		t.Errorf("Failed to set value in cache: %v", err)
	}

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value matches the expected value
	expectedValue := "value"
	if value != expectedValue {
		t.Errorf("Expected value %s, but got %s", expectedValue, value)
	}
}

func TestRedisCache_SetWithExpire(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := core_cache.NewRedisCache(client)

	// Set a value in the cache with an expiry time of 1 second
	err := cache.SetWithExpire("key", "value", 1)
	if err != nil {
		t.Errorf("Failed to set value in cache with expiry: %v", err)
	}

	// Wait for the value to expire
	// Sleep for 2 seconds to ensure the value has expired
	time.Sleep(2 * time.Second)

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value is empty
	if value != "" {
		t.Errorf("Expected value to be empty, but got %s", value)
	}
}

func TestRedisCache_Delete(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := core_cache.NewRedisCache(client)

	// Set a value in the cache
	err := cache.Set("key", "value")
	if err != nil {
		t.Errorf("Failed to set value in cache: %v", err)
	}

	// Delete the value from the cache
	err = cache.Delete("key")
	if err != nil {
		t.Errorf("Failed to delete value from cache: %v", err)
	}

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value is empty
	if value != "" {
		t.Errorf("Expected value to be empty, but got %s", value)
	}
}

func TestRedisCache_Flush(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := core_cache.NewRedisCache(client)

	// Set a value in the cache
	err := cache.Set("key", "value")
	if err != nil {
		t.Errorf("Failed to set value in cache: %v", err)
	}

	// Flush all keys in the cache
	err = cache.Flush()
	if err != nil {
		t.Errorf("Failed to flush cache: %v", err)
	}

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value is empty
	if value != "" {
		t.Errorf("Expected value to be empty, but got %s", value)
	}
}
