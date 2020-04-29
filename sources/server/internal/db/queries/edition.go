package queries

const (
	EditionGetEdition = `
		SELECT
			edition_id,
			name
		FROM
			editions
		WHERE
			edition_id = ?
	`

	EditionMarkEditionNeedPopularityRecalc = `
		UPDATE
			editions
		SET
			popularity_need_recalc = 1
		WHERE
			edition_id = ?
	`
)
