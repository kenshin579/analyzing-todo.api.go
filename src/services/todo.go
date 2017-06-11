package services

import (
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
func (svc TodoService) Add(todo types.Todo) types.Todo {
	todo.ID = uuid.NewV1().String()
	todo = svc.cacheService.Set(todo.ID, todo)
	return todo
}

// Get Todo object
func (svc TodoService) Get(id string) types.Todo {
	return svc.cacheService.Get(id)
}

// Delete Todo object
func (svc TodoService) Delete(id string) {
	svc.cacheService.Delete(id)
	return
}
