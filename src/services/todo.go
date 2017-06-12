package services

import (
	"github.com/rctyler/todo.api.go/src/shared/errors"
	"github.com/rctyler/todo.api.go/src/shared/interfaces"
	"github.com/rctyler/todo.api.go/src/shared/types"
	"github.com/satori/go.uuid"
)

// TodoService provides functionality to get, create, and delete TODO objects
type TodoService struct {
	cacheService interfaces.ICacheService
}

// NewTodoService creates a new Todo service
func NewTodoService(cacheService interfaces.ICacheService) *TodoService {
	todoService := new(TodoService)
	todoService.cacheService = cacheService
	return todoService
}

// Add a new TODO object
func (svc TodoService) Add(todo types.Todo) (types.Todo, *customErrors.Error) {
	todo.ID = uuid.NewV1().String()

	obj, err := svc.cacheService.Set(todo.ID, todo)

	todo = obj.(types.Todo)
	if e, ok := err.(*customErrors.Error); ok {
		return todo, e
	}

	return todo, nil
}

// Get TODO object
func (svc TodoService) Get(id string) (types.Todo, *customErrors.Error) {
	obj, err := svc.cacheService.Get(id)

	todo := obj.(types.Todo)
	if e, ok := err.(*customErrors.Error); ok {
		return todo, e
	}

	return todo, nil
}

// Delete TODO object
func (svc TodoService) Delete(id string) *customErrors.Error {
	err := svc.cacheService.Delete(id)

	if e, ok := err.(*customErrors.Error); ok {
		return e
	}

	return nil
}
