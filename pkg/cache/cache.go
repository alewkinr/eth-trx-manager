package cache

// Cache — cache store interface
type Cache interface {
	Get(key string) (any, bool)
	Add(key string, value any)
}
