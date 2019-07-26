package db

import "fantlab/utils"

func (db *DB) WorkExists(workId uint64) error {
	const workExistsQuery = "SELECT EXISTS(SELECT 1 FROM works WHERE work_id = ?)"

	var workExists bool

	err := db.X.Get(&workExists, workExistsQuery, workId)

	if err != nil {
		return err
	}

	if !workExists {
		return utils.ErrRecordNotFound
	}

	return nil
}
