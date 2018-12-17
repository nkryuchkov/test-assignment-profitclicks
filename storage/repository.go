package storage

import (
	"github.com/pkg/errors"
)

// AddNumberToList checks if the number list with given UID exists and, if it does, adds a new number to it.
func (s *Storage) AddNumberToList(uid string, number int64) error {
	_, err := s.getConnection().Exec("INSERT INTO `numbers` (`list_uid`, `value`) VALUES (?, ?)", uid, number)
	return errors.Wrapf(err, "could not add a new number")
}

// AddNumberList inserts a new number list with given UID in the lists table.
func (s *Storage) AddNumberList(uid string) error {
	_, err := s.getConnection().Exec("INSERT INTO `lists` (`uid`) VALUES (?)", uid)
	return err
}

// DeleteNumberList deletes a number list with given UID from the lists table.
func (s *Storage) DeleteNumberList(uid string) error {
	_, err := s.getConnection().Exec("DELETE FROM `lists` WHERE `uid` = (?)", uid)
	return err
}

// DoesNumberListExist checks if a number list with given UID exists.
func (s *Storage) DoesNumberListExist(uid string) (bool, error) {
	var cnt int
	err := s.getConnection().Get(&cnt, "SELECT COUNT(`uid`) FROM `lists` WHERE `uid` = (?)", uid)
	return cnt > 0, err
}

// AddOperationToList adds an operation to the list with given UID.
func (s *Storage) AddOperationToList(uid, operation string) error {
	_, err := s.getConnection().Exec("UPDATE `lists` SET `operation` = (?) WHERE `uid` = (?)", operation, uid)
	return err
}

// GetListOperation returns an operation that is set for the list with given UID.
func (s *Storage) GetListOperation(uid string) (string, error) {
	var operation string
	err := s.getConnection().Get(&operation, "SELECT `operation` FROM `lists` WHERE `uid` = (?)", uid)
	return operation, err
}

// GetListNumbers gets the list of numbers that relate to the number list with given UID.
func (s *Storage) GetListNumbers(uid string) ([]int64, error) {
	var numbers []int64
	err := s.getConnection().Select(&numbers, "SELECT `value` FROM `numbers` WHERE `list_uid` = (?)", uid)
	return numbers, err
}
