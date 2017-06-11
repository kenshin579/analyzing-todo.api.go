package interfaces

import (
	"github.com/rctyler/todoapp/src/shared/types"
)

// ICacheService is an interface for a caching service
type ICacheService interface {
	Get(key string) types.Todo
	Delete(key string)
	Set(key string, todo types.Todo) types.Todo
}
