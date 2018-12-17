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
	err := s.getConnection().Get(&cnt, "SELECT COUNT(uid) FROM `lists` WHERE `uid` = (?)", uid)
	return cnt > 0, err
}
