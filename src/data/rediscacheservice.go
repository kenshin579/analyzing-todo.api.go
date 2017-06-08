package data

// RedisCacheService implements ICacheService and communitcates with Redis
type RedisCacheService struct {
}

// NewRedisService creates a new RedisCache service
func NewRedisCacheService() *RedisCacheService {
	return new(RedisCacheService)
}

// Get object from Redis
func (svc RedisCacheService) Get() {

}

// Set object in Redis
func (svc RedisCacheService) Set() {

}
