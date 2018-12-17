package storage

var dbSchemaQueries = []string{
	`CREATE TABLE IF NOT EXISTS ` + `lists` + ` (
	` + `uid` + ` varchar(255) NOT NULL,
	` + `operation` + ` varchar(255),
	PRIMARY KEY (` + `uid` + `));`,

	`CREATE TABLE IF NOT EXISTS ` + `numbers` + ` (
	` + `id` + ` INT NOT NULL AUTO_INCREMENT,
	` + `list_uid` + ` varchar(255) NOT NULL,
	` + `value` + ` INT NOT NULL,
	PRIMARY KEY (` + `id` + `),
	CONSTRAINT ` + `numbers_fk0` + ` FOREIGN KEY (` + `list_uid` + `) REFERENCES ` +
		`lists` + `(` + `uid` + `));`,

	`CREATE TABLE IF NOT EXISTS ` + `logs` + ` (
	` + `id` + ` INT NOT NULL AUTO_INCREMENT,
	` + `text` + ` varchar(255) NOT NULL,
	PRIMARY KEY (` + `id` + `));`,
}

// CreateSchemaIfNotExists creates a database schema if it doesn't exist
func (s *Storage) CreateSchemaIfNotExists() error {
	for _, query := range dbSchemaQueries {
		if _, err := s.getConnection().Exec(query); err != nil {
			return err
		}
	}
	return nil
}
