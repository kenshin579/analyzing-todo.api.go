package services

import "github.com/rctyler/todoapp/src/common/interfaces"

// TodoService provides functionality to get, create, update/patch, and delete todo reminders
type TodoService struct {
	cacheService services.ICacheService
}

// Todo is a struct that describes a TODO reminder
type Todo struct {
	ID          string
	IsCompleted bool
}

// NewTodoService creates a new Todo service
func NewTodoService(cacheService services.ICacheService) *TodoService {
	todoService := new(TodoService)
	todoService.cacheService = cacheService
	return todoService
}

// GetAll Todo objects
func (svc TodoService) GetAll() []Todo {
	message := svc.cacheService.Get()

	return []Todo{
		Todo{ID: "12345", IsCompleted: true},
		Todo{ID: message, IsCompleted: false}}
}

// Add Todo objects
func (svc TodoService) Add() Todo {
	return Todo{ID: "12345", IsCompleted: true}
}

// Get Todo object
func (svc TodoService) Get(id string) Todo {
	return Todo{ID: id, IsCompleted: true}
}

// Update Todo objects
func (svc TodoService) Update(id string) Todo {
	return Todo{ID: id, IsCompleted: true}
}

// Patch Todo objects
func (svc TodoService) Patch(id string) Todo {
	return Todo{ID: id, IsCompleted: true}
}

// Delete Todo objects
func (svc TodoService) Delete(id string) Todo {
	return Todo{ID: id, IsCompleted: true}
}
