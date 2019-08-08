package db

func (db *DB) WorkExists(workId uint64) error {
	var workExists uint8
	err := db.R.Query("SELECT 1 FROM works WHERE work_id = ?", workId).Scan(&workExists)
	return err
}
