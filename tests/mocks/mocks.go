package mocks

import "github.com/rctyler/todo.api.go/src/shared/types"

// MockCacheService mocks ICacheService and communitcates with Redis
type MockCacheService struct {
}

// NewMockCacheService creates a new RedisCache service
func NewMockCacheService() *MockCacheService {
	mockCacheService := new(MockCacheService)
	return mockCacheService
}

// Get object from Redis
func (svc MockCacheService) Get(key string) (interface{}, error) {
	return types.Todo{}, nil
}

// Delete object in Redis
func (svc MockCacheService) Delete(key string) error {
	return nil
}

// Set object in Redis
func (svc MockCacheService) Set(key string, obj interface{}) (interface{}, error) {
	return types.Todo{}, nil
}
