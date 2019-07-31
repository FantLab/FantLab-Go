package db

func (db *DB) WorkExists(workId uint64) error {
	var workExists bool
	err := db.X.Get(&workExists, "SELECT 1 FROM works WHERE work_id = ?", workId)
	return err
}
