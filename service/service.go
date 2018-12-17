package service

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/nkryuchkov/test-assignment-profitclicks/logger"
	"github.com/nkryuchkov/test-assignment-profitclicks/storage"
)

var (
	ErrListNotExists      = errors.New("list with this UID does not exist")
	ErrOperationNotExists = errors.New("operation with this name does not exist")
)

// Service represents a structure which methods provide the app business logic.
type Service struct {
	log     *logger.Logger
	storage *storage.Storage
}

// New returns a new service instance.
func New(log *logger.Logger, storage *storage.Storage) *Service {
	return &Service{
		log:     log,
		storage: storage,
	}
}

// AddNumberToList adds a new number to the list with given UID.
func (s *Service) AddNumberToList(uid string, number int64) error {
	exist, err := s.storage.DoesNumberListExist(uid)
	if err != nil {
		return errors.Wrapf(err, "could not check if number list exists")
	}
	if !exist {
		return ErrListNotExists
	}

	if err = s.storage.AddNumberToList(uid, number); err != nil {
		return err
	}

	return nil
}

// AddNumberList adds a new number list to storage.
func (s *Service) AddNumberList() (string, error) {
	uid := s.generateUID()
	err := s.storage.AddNumberList(uid)
	return uid, err
}

// DeleteNumberList deletes a number list from storage.
func (s *Service) DeleteNumberList(uid string) error {
	exist, err := s.storage.DoesNumberListExist(uid)
	if err != nil {
		return errors.Wrapf(err, "could not check if number list exists")
	}
	if !exist {
		return ErrListNotExists
	}

	return s.storage.DeleteNumberList(uid)
}

// AddOperationToList adds an operation to the list.
func (s *Service) AddOperationToList(uid, operation string) error {
	exist, err := s.storage.DoesNumberListExist(uid)
	if err != nil {
		return errors.Wrapf(err, "could not check if number list exists")
	}
	if !exist {
		return ErrListNotExists
	}

	if !s.isAllowedOperation(operation) {
		return ErrOperationNotExists
	}

	return s.storage.AddOperationToList(uid, operation)
}

// GetListResult fetches an operation for the list with given UID, fetches a list of numbers in that list,
// performs the operation on the list and returns its result.
func (s *Service) GetListResult(uid string) (int64, error) {
	exist, err := s.storage.DoesNumberListExist(uid)
	if err != nil {
		return 0, errors.Wrapf(err, "could not check if number list exists")
	}
	if !exist {
		return 0, ErrListNotExists
	}

	operation, err := s.storage.GetListOperation(uid)
	if err != nil {
		return 0, errors.Wrapf(err, "could not get list operation")
	}

	numbers, err := s.storage.GetListNumbers(uid)
	if err != nil {
		return 0, errors.Wrapf(err, "could not get list numbers")
	}

	return s.mapOperationToFunc(operation)(numbers), nil
}

func (s *Service) generateUID() string {
	t := time.Now().UnixNano()
	str := strconv.FormatInt(t, 10)
	uid := fmt.Sprintf("%x", sha1.Sum([]byte(str)))

	return uid
}
