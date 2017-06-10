package data

import "github.com/go-redis/redis"

// RedisCacheService implements ICacheService and communitcates with Redis
type RedisCacheService struct {
	client *redis.Client
}

// NewRedisService creates a new RedisCache service
func NewRedisCacheService(address string, password string, database int) *RedisCacheService {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       database,
	})

	redisCacheService := new(RedisCacheService)
	redisCacheService.client = client

	return redisCacheService
}

// Get object from Redis
func (svc RedisCacheService) Get() string {
	pong, _ := svc.client.Ping().Result()
	return pong
}

// Set object in Redis
func (svc RedisCacheService) Set() {

}
