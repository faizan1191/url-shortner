package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	Save(code, url string) error
	Get(code string) (string, bool)
}

type RedisStore struct {
	client *redis.Client   // Redis connection
	ctx    context.Context // context for Redis commands
}

func NewRedisStore(addr string) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr, // e.g. "localhost:6379"
	})
	// addr is Redis server address
	return &RedisStore{
		client: rdb,
		ctx:    context.Background(), // creates a base context for redis commands
	}
}

// Saves key-value pair in redis and returns error if any
func (r *RedisStore) Save(code, url string) error {
	return r.client.Set(r.ctx, code, url, 0).Err()
}

// Get value from key in redis store

func (r *RedisStore) Get(code string) (string, bool) {
	val, err := r.client.Get(r.ctx, code).Result()

	if err == redis.Nil || err != nil {
		return "", false
	}

	return val, true
}
