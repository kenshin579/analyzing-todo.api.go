package services

// ICacheService is an interface for a caching service
type ICacheService interface {
	Get() string
	Set()
}
