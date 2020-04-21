package queries

const (
	BookcasesTable = "bookcase"
)

const (
	BookcaseGetBookcases = `
		SELECT
			b.bookcase_id,
			b.user_id,
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

	BookcaseGetBookcase = `
		SELECT
			bookcase_id,
			user_id,
			bookcase_type,
			bookcase_group,
			bookcase_name,
			bookcase_comment,
			bookcase_shared,
			sort
		FROM
			bookcase
		WHERE
			bookcase_type = ? AND bookcase_id = ?
		LIMIT 1
	`

	BookcaseGetFilmBookcaseItems = `
		SELECT
			bi.item_id AS 'film_id',
			bi.item_comment AS 'comment',
			f.name,
			f.rusname,
			f.type,
			f.year,
			f.year2,
			f.country,
			f.genre,
			f.director,
			f.screenwriter,
			f.actors,
			f.description
		FROM
			bookcase_items bi
		LEFT JOIN
			films f ON f.film_id = bi.item_id
		WHERE
			bi.bookcase_id = ?
		ORDER BY
			%s
		LIMIT ?
		OFFSET ?
	`

	BookcaseGetFilmBookcaseItemCount = `
		SELECT
			COUNT(*)
		FROM
			bookcase_items
		WHERE
			bookcase_id = ?
	`
)
