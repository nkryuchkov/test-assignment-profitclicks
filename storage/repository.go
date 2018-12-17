package storage

// AddNumberList inserts a new number list with given uid in the lists table.
func (s *Storage) AddNumberList(uid string) error {
	_, err := s.getConnection().Exec("INSERT INTO `lists` (uid) VALUES (?)", uid)
	return err
}

// DeleteNumberList deletes a number list with given uid from the lists table.
func (s *Storage) DeleteNumberList(uid string) error {
	_, err := s.getConnection().Exec("DELETE FROM `lists` WHERE uid = (?)", uid)
	return err
}
