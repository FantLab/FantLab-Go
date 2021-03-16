package queries

const (
	AutorGetAutors = `
		SELECT
			autor_id,
			shortrusname
		FROM
			autors
		WHERE
			autor_id IN (?)
	`

	AutorIncrementAutorsNewResponseCount = `
		UPDATE
			autors
		SET
			new_responses_count = new_responses_count + 1
		WHERE
			autor_id IN (?)
	`

	AutorDecrementAutorsNewResponseCount = `
		UPDATE
			autors
		SET
			new_responses_count = new_responses_count - 1
		WHERE
			autor_id IN (?) AND new_responses_count > 0
	`

	AutorMarkAutorsNeedRecalcStats = `
		UPDATE
			autor_stats
		SET
			is_actual = 0
		WHERE
			autor_id IN (?)
	`
)
