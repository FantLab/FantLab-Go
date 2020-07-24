package queries

const (
	WorkStatsInsertResponseCount = `
		INSERT INTO
			work_stats (
				work_id,
				responsecount
			)
		VALUES
			(?, 1)
		ON DUPLICATE KEY UPDATE
			responsecount = responsecount + 1
	`

	// NOTE Неясно, что представляет собой поле responsecount1 и почему оно изменяется при удалении отзыва, но не
	// трогается при добавлении
	WorkStatsDecrementResponseCount = `
		UPDATE
			work_stats
		SET
			responsecount = responsecount - 1,
			responsecount1 = responsecount1 - 1
		WHERE
			work_id = ?
	`

	WorkStatsUpsertWorkStats = `
		INSERT INTO
			work_stats (
				work_id,
				midmark,
				midmark_by_weight,
				markcount,
				rating,
				gendermidmarkdelta,
				midmark_male,
				midmark_female,
				markcount_male,
				markcount_female,
				markmiddledelta,
				date_of_calc
			)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())
		ON DUPLICATE KEY UPDATE
			midmark = ?,
			midmark_by_weight = ?,
			markcount = ?,
			rating = ?,
			gendermidmarkdelta = ?,
			midmark_male = ?,
			midmark_female = ?,
			markcount_male = ?,
			markcount_female = ?,
			markmiddledelta = ?,
			date_of_calc = NOW()
	`

	WorkStatsDeleteWorkStats = `
		DELETE FROM
			work_stats
		WHERE
			work_id = ?
	`
)
