package queries

const (
	ResponseDeleteResponsePreview = `
		DELETE
		FROM
			responses_preview
		WHERE
			user_id = ? AND work_id = ?
	`
)
