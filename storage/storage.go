package storage

import (
	"github.com/nkryuchkov/test-assignment-profitclicks/logger"
)

// Config represents a storage configuration.
type Config struct {
}

// Storage is a structure with methods that interact with data storage.
type Storage struct {
	config *Config
	log    *logger.Logger
}

// New returns a new storage.
func New(config *Config, log *logger.Logger) *Storage {
	return &Storage{
		config: config,
		log:    log,
	}
}
