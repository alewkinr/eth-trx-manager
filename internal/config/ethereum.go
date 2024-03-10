package config

// Ethereum — ethereum specific settings
type Ethereum struct {
	// URL — URL of the Ethereum node
	URL string `koanf:"url"`

	PrivateKey string `koanf:"private_key"`
}
