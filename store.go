package app

import "github.com/go-redis/redis"

type Store interface {
	Set(key, value string) error
	Get(key string) (string, error)
}

type redisStore struct {
	client *redis.Client
}

func newRedisStore(c *redis.Client) Store {
	return &redisStore{client: c}
}

func (s *redisStore) Set(key, value string) error {
	return s.client.Set(key, value, 0).Err()
}

func (s *redisStore) Get(key string) (string, error) {
	return s.client.Get(key).Result()
}
