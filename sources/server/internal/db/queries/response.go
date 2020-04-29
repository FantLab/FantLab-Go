package queries

const (
	ResponseGetResponse = `
		SELECT
			response_id,
			user_id
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
)
