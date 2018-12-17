package storage

import (
	"sync"

	// Using a mysql driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nkryuchkov/test-assignment-profitclicks/logger"
	"github.com/pkg/errors"
)

// Config represents a storage configuration.
type Config struct {
	DriverName string `json:"driver_name"`
	Connection string `json:"connection"`
}

// Storage is a structure with methods that interact with data storage.
type Storage struct {
	config *Config
	log    *logger.Logger
	sync.Mutex
	conn *sqlx.DB
}

// New returns a new storage instance.
func New(config *Config, log *logger.Logger) *Storage {
	return &Storage{
		config: config,
		log:    log,
	}
}

// Close closes the database instance.
func (s *Storage) Close() error {
	return s.getConnection().Close()
}

// Connect connects to the storage instance.
func (s *Storage) Connect() error {
	conn, err := sqlx.Connect(s.config.DriverName, s.config.Connection)
	if err != nil {
		return errors.Wrapf(err, "could not connect to %v", s.config.Connection)
	}

	s.setConnection(conn)
	return nil
}

func (s *Storage) getConnection() *sqlx.DB {
	s.Lock()
	defer s.Unlock()

	return s.conn
}

func (s *Storage) setConnection(conn *sqlx.DB) {
	s.Lock()
	defer s.Unlock()

	s.conn = conn
}
