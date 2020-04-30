package queries

const (
	AutorDecrementAutorsNewResponseCount = `
		UPDATE
			autors
		SET
			new_responses_count = new_responses_count - 1
		WHERE
			autor_id IN (?)
	`
)
