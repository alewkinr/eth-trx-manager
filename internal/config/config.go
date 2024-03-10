package config

import (
	"fmt"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	k *koanf.Koanf `koanf:"-"`

	// HTTP server settings
	// Host — http service host
	Host string `koanf:"host"`
	// Port — http service port
	Port string `koanf:"port"`

	*Ethereum `koanf:"ethereum"`
}

const (
	delim     = "_"
	envPrefix = "ETM_"
)

// MustNewConfig — constructor for configuration struct, or panic if error
func MustNewConfig() *Config {
	cfg := &Config{
		k: koanf.New("."),
	}

	if err := cfg.k.Load(env.Provider("", "_", nil), nil); err != nil {
		panic(fmt.Errorf("providing env vars: %w", err))
	}

	if unmarshalErr := cfg.k.Unmarshal("", cfg); unmarshalErr != nil {
		panic(fmt.Errorf("unmarshaling config: %w", unmarshalErr))
	}

	return cfg
}
