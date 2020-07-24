package queries

const (
	CorrelationsInsertUserWorkUpdate = `
		INSERT INTO
			korel_work_update (
				user_id,
				work_id
			)
		VALUES
			(?, ?)
		ON DUPLICATE KEY UPDATE
			flag = 0
	`
)
