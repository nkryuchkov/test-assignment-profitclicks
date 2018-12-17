package storage

// AddNumberList inserts a new number list with given uid in the lists table.
func (s *Storage) AddNumberList(uid string) error {
	_, err := s.getConnection().Exec(`INSERT INTO lists (uid) VALUES(?)`, uid)
	return err
}
