package queries

const (
	BookcasesTable = "bookcase"
)

const (
	BookcaseGetBookcases = `
		SELECT
			b.bookcase_id,
			b.bookcase_type,
			b.bookcase_group,
			b.bookcase_name,
			b.bookcase_comment,
			b.bookcase_shared,
			b.sort,
			COUNT(bi.bookcase_id) AS item_count
		FROM
			bookcase b
		LEFT JOIN
			bookcase_items bi ON bi.bookcase_id = b.bookcase_id
		WHERE
			b.user_id = ? AND %s
		GROUP BY
			b.bookcase_id
		ORDER BY
			b.bookcase_type,
			b.sort,
			b.date_of_add
	`
)
