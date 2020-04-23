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

	BookcaseGetEditionBookcaseItems = `
		SELECT
			e.edition_id,
			e.name,
			e.autors,
			e.type,
			e.year,
			e.publisher,
			e.description,
			e.correct,
			e.plan_date,
			e.ozon_id,
			e.ozon_cost,
			e.ozon_available,
			e.labirint_id,
			e.labirint_cost,
			e.labirint_available,
			bi.item_comment AS 'comment'
		FROM
			bookcase_items bi
		LEFT JOIN
			editions e ON e.edition_id = bi.item_id
		WHERE
			bi.bookcase_id = ?
		ORDER BY
			%s
		LIMIT ?
		OFFSET ?
	`

	BookcaseGetFilmBookcaseItems = `
		SELECT
			f.film_id,
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
			f.description,
			bi.item_comment AS 'comment'
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

	BookcaseGetBookcaseItemCount = `
		SELECT
			COUNT(*)
		FROM
			bookcase_items
		WHERE
			bookcase_id = ?
	`
)
