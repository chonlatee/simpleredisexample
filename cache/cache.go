package cache

import "github.com/go-redis/redis"

// Cache struct ...
type Cache struct {
	redisCache *redis.Client
}

func connectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.2:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}

// Set cache Set(key, val string) error
func (c *Cache) Set(key, val string) error {
	err := c.redisCache.Set(key, val, 0).Err()
	return err
}

// Get cache Get(key string) string
func (c *Cache) Get(key string) (string, error) {
	val, err := c.redisCache.Get(key).Result()
	return val, err
}

// NewCache for new cache system
func NewCache() *Cache {
	rdb := connectRedis()

	return &Cache{
		redisCache: rdb,
	}
}
