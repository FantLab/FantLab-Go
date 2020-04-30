package queries

const (
	ResponseGetResponse = `
		SELECT
			response_id,
			user_id,
			work_id
		FROM
			responses
		WHERE
			response_id = ?
	`

	ResponseUpdateResponse = `
		UPDATE
			responses
		SET
			response = ?
		WHERE
			response_id = ?
	`

	ResponseDeleteResponse = `
		DELETE
		FROM
			responses
		WHERE
			response_id = ?
	`
)
