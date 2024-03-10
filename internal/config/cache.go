package config

import "time"

// Cache — cache settings
type Cache struct {
	Size int           `koanf:"size"`
	TTL  time.Duration `koanf:"TTL"`
}
