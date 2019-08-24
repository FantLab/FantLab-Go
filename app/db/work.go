package db

import (
	"fantlab/dbtools/sqlr"
)

var workExistsQuery = sqlr.NewQuery("SELECT 1 FROM works WHERE work_id = ?")

func (db *DB) WorkExists(workId uint64) (bool, error) {
	var workExists uint8
	err := db.R.Read(workExistsQuery.WithArgs(workId)).Scan(&workExists)
	return workExists == 1, err
}
