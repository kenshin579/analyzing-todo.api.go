package data

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/rctyler/todoapp/src/data/daos"
	"github.com/rctyler/todoapp/src/shared/errors"
	"github.com/rctyler/todoapp/src/shared/types"
)

// RedisCacheService implements ICacheService and communitcates with Redis
type RedisCacheService struct {
	client *redis.Client
}

// NewRedisCacheService creates a new RedisCache service
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
func (svc RedisCacheService) Get(key string) (types.Todo, *customErrors.Error) {
	var todoDao daos.Todo

	res, err := svc.client.Get(key).Result()
	if err == redis.Nil {
		return types.Todo{}, customErrors.New("NOT_FOUND", "Requested TODO not found")
	} else if err != nil {
		return types.Todo{}, customErrors.New("SERVICE_UNAVAILABLE", err.Error())
	}

	if err := json.Unmarshal([]byte(res), &todoDao); err != nil {
		return types.Todo{}, customErrors.New("INTERNAL_SERVER_ERROR", err.Error())
	}

	todo := types.Todo{
		ID:     todoDao.ID,
		TODO:   todoDao.TODO,
		Author: todoDao.Author,
		When:   todoDao.When,
	}

	return todo, nil
}

// Delete object from Redis
func (svc RedisCacheService) Delete(key string) *customErrors.Error {
	_, err := svc.client.Del(key).Result()
	if err != nil {
		return customErrors.New("SERVICE_UNAVAILABLE", err.Error())
	}

	return nil
}

// Set object in Redis
func (svc RedisCacheService) Set(key string, obj types.Todo) (types.Todo, *customErrors.Error) {
	todoDao := daos.Todo{
		ID:     obj.ID,
		TODO:   obj.TODO,
		Author: obj.Author,
		When:   obj.When,
	}

	todoJSON, err := json.Marshal(todoDao)
	if err != nil {
		return types.Todo{}, customErrors.New("INTERNAL_SERVER_ERROR", err.Error())
	}

	_, err = svc.client.Set(key, string(todoJSON), 0).Result()
	if err != nil {
		return types.Todo{}, customErrors.New("SERVICE_UNAVAILABLE", err.Error())
	}

	return obj, nil
}
