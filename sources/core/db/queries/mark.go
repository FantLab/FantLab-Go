package queries

const (
	MarkGetWorkMarks = `
		SELECT
			m.mark,
			u.mark_weight,
			u.sex
		FROM
			marks2 m
		JOIN
			users u ON (u.user_id = m.user_id)
		WHERE
			m.work_id = ?
	`

	MarkUpsertUserWorkMark = `
		INSERT INTO
			marks2 (
				user_id,
				work_id,
				mark,
				date_of_add,
				date_of_last_change
			)
		VALUES
			(?, ?, ?, NOW(), NOW())
		ON DUPLICATE KEY UPDATE
			mark = ?,
			date_of_last_change = NOW()
	`

	MarkDeleteUserWorkMark = `
		DELETE FROM
			marks2
		WHERE
			user_id = ? AND work_id = ?
	`
)
