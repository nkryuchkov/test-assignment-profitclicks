package service

import (
	"github.com/nkryuchkov/test-assignment-profitclicks/logger"
)

// Service represents a structure which methods provide the app business logic.
type Service struct {
	log *logger.Logger
}

// New returns a new service instance.
func New(log *logger.Logger) *Service {
	return &Service{
		log: log,
	}
}

func (s *Service) AddNumberToList(listID string, number int64) error {
	return nil
}

func (s *Service) AddNumberList() (string, error) {
	return "", nil
}

func (s *Service) DeleteNumberList(listID string) error {
	return nil
}

func (s *Service) AddOperationToList(listID string, operationName int64) error {
	return nil
}

func (s *Service) GetListResult(listID string) (int, error) {
	return 0, nil
}
