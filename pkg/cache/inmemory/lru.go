package inmemory

import (
	"log/slog"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

type Store struct {
	cache *expirable.LRU[string, any]
	log   *slog.Logger
}

// NewStore — create new inmemory store
func NewStore(size int, ttl time.Duration, logger *slog.Logger) *Store {
	logger.Debug("💿 Cache initialized", "size", size, "ttl", ttl.String())

	return &Store{
		cache: expirable.NewLRU[string, any](size, nil, ttl),
		log:   logger,
	}
}

func (s *Store) Get(key string) (any, bool) {
	value, ok := s.cache.Get(key)

	s.log.Debug("Cache.Get", "found", ok, "key", key, "value", value)

	return value, ok
}

func (s *Store) Add(key string, value any) {
	s.log.Debug("Cache.Add", "key", key, "value", value)
	s.cache.Add(key, value)
}
