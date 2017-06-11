package services

import (
	"github.com/rctyler/todoapp/src/shared/errors"
	"github.com/rctyler/todoapp/src/shared/interfaces"
	"github.com/rctyler/todoapp/src/shared/types"
	"github.com/satori/go.uuid"
)

// TodoService provides functionality to get, create, update/patch, and delete todo reminders
type TodoService struct {
	cacheService interfaces.ICacheService
}

// NewTodoService creates a new Todo service
func NewTodoService(cacheService interfaces.ICacheService) *TodoService {
	todoService := new(TodoService)
	todoService.cacheService = cacheService
	return todoService
}

// Add Todo objectstypes.
func (svc TodoService) Add(todo types.Todo) (types.Todo, *customErrors.Error) {
	todo.ID = uuid.NewV1().String()
	return svc.cacheService.Set(todo.ID, todo)
}

// Get Todo object
func (svc TodoService) Get(id string) (types.Todo, *customErrors.Error) {
	return svc.cacheService.Get(id)
}

// Delete Todo object
func (svc TodoService) Delete(id string) *customErrors.Error {
	return svc.cacheService.Delete(id)
}
