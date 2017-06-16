package mocks

import (
	"github.com/rctyler/todo.api.go/src/shared/errors"
	"github.com/rctyler/todo.api.go/src/shared/types"
)

// MockCacheService mocks ICacheService and communitcates with Redis
type MockCacheService struct {
	TestOptions TestOptions
}

// TestOptions setup MockCacheService
type TestOptions struct {
	ShouldReturnError bool
}

// NewMockCacheService creates a new RedisCache service
func NewMockCacheService(testOptions TestOptions) *MockCacheService {
	mockCacheService := new(MockCacheService)
	mockCacheService.TestOptions = testOptions
	return mockCacheService
}

// Get object from Redis
func (mock MockCacheService) Get(key string) (interface{}, error) {
	if mock.TestOptions.ShouldReturnError {
		return types.Todo{}, customErrors.New("ERROR", "ERROR: MockCacheService.Get")
	}
	return types.Todo{}, nil
}

// Delete object in Redis
func (mock MockCacheService) Delete(key string) error {
	if mock.TestOptions.ShouldReturnError {
		return customErrors.New("ERROR", "ERROR: MockCacheService.Delete")
	}
	return nil
}

// Set object in Redis
func (mock MockCacheService) Set(key string, obj interface{}) (interface{}, error) {
	if mock.TestOptions.ShouldReturnError {
		return types.Todo{}, customErrors.New("ERROR", "ERROR: MockCacheService.Set")
	}
	return types.Todo{}, nil
}
