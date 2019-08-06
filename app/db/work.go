package db

func (db *DB) WorkExists(workId uint64) error {
	var workExists uint8
	err := db.R.Query(&workExists, "SELECT 1 FROM works WHERE work_id = ?", workId)
	return err
}
