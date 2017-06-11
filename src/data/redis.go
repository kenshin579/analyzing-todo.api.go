package data

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/rctyler/todoapp/src/data/daos"
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
func (svc RedisCacheService) Get(key string) types.Todo {
	var todoDao daos.Todo

	res, err := svc.client.Get(key).Result()
	if err == redis.Nil {
		return types.Todo{}
	} else if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(res), &todoDao)
	if err != nil {
		panic(err)
	}

	todo := types.Todo{
		ID:     todoDao.ID,
		TODO:   todoDao.TODO,
		Author: todoDao.Author,
		When:   todoDao.When,
	}

	return todo
}

// Delete object from Redis
func (svc RedisCacheService) Delete(key string) {
	_, err := svc.client.Del(key).Result()
	if err != nil {
		panic(err)
	}

	return
}

// Set object in Redis
func (svc RedisCacheService) Set(key string, obj types.Todo) types.Todo {
	todoDao := daos.Todo{
		ID:     obj.ID,
		TODO:   obj.TODO,
		Author: obj.Author,
		When:   obj.When,
	}

	b, err := json.Marshal(todoDao)
	if err != nil {
		panic(err)
	}

	_, err = svc.client.Set(key, string(b), 0).Result()
	if err != nil {
		panic(err)
	}

	return obj
}
