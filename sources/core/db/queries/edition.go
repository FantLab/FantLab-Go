package queries

const (
	EditionGetEditions = `
		SELECT
			edition_id,
			name
		FROM
			editions
		WHERE
			edition_id IN (?)
	`

	EditionMarkEditionsNeedPopularityRecalc = `
		UPDATE
			editions
		SET
			popularity_need_recalc = 1
		WHERE
			edition_id IN (?)
	`
)
