package todo_test

import (
	"testing"

	"github.com/rctyler/todo.api.go/src/services"
	"github.com/rctyler/todo.api.go/src/shared/types"
	"github.com/rctyler/todo.api.go/tests/mocks"
)

func TestGet(t *testing.T) {
	mockCacheService := mocks.NewMockCacheService()
	todoService := services.NewTodoService(mockCacheService)

	id := "string"

	_, err := todoService.Get(id)
	if err != nil {
		t.Error()
	}
}

func TestAdd(t *testing.T) {
	mockCacheService := mocks.NewMockCacheService()
	todoService := services.NewTodoService(mockCacheService)

	todo := types.Todo{
		TODO:   "string",
		Author: "string",
		When:   "string",
	}

	_, err := todoService.Add(todo)
	if err != nil {
		t.Error()
	}
}

func TestDelete(t *testing.T) {
	mockCacheService := mocks.NewMockCacheService()
	todoService := services.NewTodoService(mockCacheService)

	id := "string"

	err := todoService.Delete(id)
	if err != nil {
		t.Error()
	}
}
