package tests

import (
	"testing"

	"github.com/rctyler/todo.api.go/src/services"
	"github.com/rctyler/todo.api.go/src/shared/types"
	"github.com/rctyler/todo.api.go/tests/mocks"
)

func TestGet(t *testing.T) {
	// ARRANGE
	mockCacheService := mocks.NewMockCacheService(mocks.TestOptions{})
	todoService := services.NewTodoService(mockCacheService)

	id := "string"

	// ACT
	_, err := todoService.Get(id)

	// ASSERT
	if err != nil {
		t.Error()
	}
}

func TestGetError(t *testing.T) {
	// ARRANGE
	mockCacheService := mocks.NewMockCacheService(mocks.TestOptions{ShouldReturnError: true})
	todoService := services.NewTodoService(mockCacheService)

	id := "string"

	// ACT
	_, err := todoService.Get(id)

	// ASSERT
	if err == nil || err.Error() != "ERROR: MockCacheService.Get" {
		t.Error()
	}
}

func TestAdd(t *testing.T) {
	// ARRANGE
	mockCacheService := mocks.NewMockCacheService(mocks.TestOptions{})
	todoService := services.NewTodoService(mockCacheService)

	todo := types.Todo{
		TODO:   "string",
		Author: "string",
		When:   "string",
	}

	// ACT
	_, err := todoService.Add(todo)

	// ASSERT
	if err != nil {
		t.Error()
	}
}

func TestAddError(t *testing.T) {
	// ARRANGE
	mockCacheService := mocks.NewMockCacheService(mocks.TestOptions{ShouldReturnError: true})
	todoService := services.NewTodoService(mockCacheService)

	todo := types.Todo{
		TODO:   "string",
		Author: "string",
		When:   "string",
	}

	// ACT
	_, err := todoService.Add(todo)

	// ASSERT
	if err == nil || err.Error() != "ERROR: MockCacheService.Set" {
		t.Error()
	}
}

func TestDelete(t *testing.T) {
	// ARRANGE
	mockCacheService := mocks.NewMockCacheService(mocks.TestOptions{})
	todoService := services.NewTodoService(mockCacheService)

	id := "string"

	// ACT
	err := todoService.Delete(id)

	// ASSERT
	if err != nil {
		t.Error()
	}
}

func TestDeleteError(t *testing.T) {
	// ARRANGE
	mockCacheService := mocks.NewMockCacheService(mocks.TestOptions{ShouldReturnError: true})
	todoService := services.NewTodoService(mockCacheService)

	id := "string"

	// ACT
	err := todoService.Delete(id)

	// ASSERT
	if err == nil || err.Error() != "ERROR: MockCacheService.Delete" {
		t.Error()
	}
}
