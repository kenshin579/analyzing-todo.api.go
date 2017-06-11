package interfaces

import (
	"github.com/rctyler/todoapp/src/shared/errors"
	"github.com/rctyler/todoapp/src/shared/types"
)

// ICacheService is an interface for a caching service
type ICacheService interface {
	Get(key string) (types.Todo, *customErrors.Error)
	Delete(key string) *customErrors.Error
	Set(key string, todo types.Todo) (types.Todo, *customErrors.Error)
}
