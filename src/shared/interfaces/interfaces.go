package interfaces

// ICacheService is an interface for a generic caching service. This doesn't need to be as generic as it is, but I wanted to play around with type conversion
type ICacheService interface {
	Get(key string) (interface{}, error)
	Delete(key string) error
	Set(key string, obj interface{}) (interface{}, error)
}
