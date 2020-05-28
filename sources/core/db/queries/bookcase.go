package queries

const (
	BookcasesTable     = "bookcase"
	BookcaseItemsTable = "bookcase_items"
)

const (
	BookcaseGetAllUserBookcases = `
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

	BookcaseGetUserBookcasesSort = `
		SELECT
			bookcase_id,
			sort
		FROM
			bookcase
		WHERE
			user_id = ? AND bookcase_id IN (?)
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
			bookcase_id = ?
		LIMIT 1
	`

	BookcaseGetTypedBookcase = `
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

	BookcaseGetBookcaseItem = `
		SELECT
			bookcase_item_id,
			bookcase_id,
			item_id
		FROM
			bookcase_items
		WHERE
			bookcase_item_id = ?
	`

	BookcaseGetEditionBookcaseItems = `
		SELECT
			bi.bookcase_item_id AS item_id,
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
			bi.item_comment AS comment
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

	BookcaseGetWorkBookcaseItems = `
		SELECT
			bi.bookcase_item_id AS item_id,
			w.work_id,
			w.name,
			w.altname,
			w.rusname,
			w.year,
			w.bonus_text,
			w.description,
			w.published,
			w.work_type_id,
			w.autor_id,
			w.autor2_id,
			w.autor3_id,
			w.autor4_id,
			w.autor5_id,
			ws.midmark,
			ws.markcount,
			COUNT(r.work_id) AS response_count,
			bi.item_comment AS comment
		FROM
			bookcase_items bi
		JOIN
			works w ON w.work_id = bi.item_id
		LEFT JOIN
			work_stats ws ON ws.work_id = w.work_id
		LEFT JOIN
			responses r ON r.work_id = w.work_id
		WHERE
			bi.bookcase_id = ?
		GROUP BY
			w.work_id
		ORDER BY
			%s
		LIMIT ?
		OFFSET ?
	`

	BookcaseGetWorksAutors = `
		SELECT
			autor_id,
			rusname,
			is_opened
		FROM
			autors
		WHERE
			autor_id IN (?)
	`

	BookcaseGetOwnWorkMarks = `
		SELECT
			work_id,
			mark
		FROM
			marks2
		WHERE
			work_id IN (?) AND user_id = ?
	`

	BookcaseGetOwnWorkResponses = `
		SELECT
			work_id,
			1
		FROM
			responses
		WHERE
			work_id IN (?) AND user_id = ?
	`

	BookcaseGetFilmBookcaseItems = `
		SELECT
			bi.bookcase_item_id AS item_id,
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
			bi.item_comment AS comment
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

	BookcaseGetBookcaseItemsDateOfAdd = `
		SELECT
			item_id,
			date_of_add
		FROM
			bookcase_items
		WHERE
			bookcase_id = ?
	`

	BookcaseGetBookcaseItemCount = `
		SELECT
			COUNT(*)
		FROM
			bookcase_items
		WHERE
			bookcase_id = ?
	`

	BookcaseGetMaxSortForType = `
		SELECT
			MAX(sort)
		FROM
			bookcase
		WHERE
			user_id = ? AND bookcase_type = ?
	`

	BookcaseInsertBookcase = `
		INSERT INTO
			bookcase (
				user_id,
				bookcase_type,
				bookcase_group,
				bookcase_name,
				bookcase_comment,
				bookcase_shared,
				sort,
				date_of_add,
				default_sort
			)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, NOW(), ?)
	`

	BookcaseInsertItem = `
		INSERT IGNORE
			bookcase_items (
				bookcase_id,
				item_id,
				item_sort,
				date_of_add
			)
		SELECT
			?, ?, COALESCE(MAX(item_sort), 0) + 1, NOW()
		FROM
			bookcase_items
		WHERE
			bookcase_id = ?
	`

	BookcaseUpdateBookcase = `
		UPDATE
			bookcase
		SET
			bookcase_name = ?,
			bookcase_comment = ?,
			bookcase_shared = ?,
			bookcase_group = ?,
			default_sort = ?,
			date_of_edit = NOW()
		WHERE
			bookcase_id = ?
	`

	BookcaseUpdateSort = `
		UPDATE
			bookcase
		SET
			sort = ?
		WHERE
			bookcase_id = ?
	`

	BookcaseUpdateItemComment = `
		UPDATE
			bookcase_items
		SET
			item_comment = ?
		WHERE
			bookcase_item_id = ?
	`

	BookcaseDeleteBookcaseItems = `
		DELETE
		FROM
			bookcase_items
		WHERE
			bookcase_id = ?
	`

	BookcaseDeleteItem = `
		DELETE
		FROM
			bookcase_items
		WHERE
			bookcase_item_id = ?
	`

	BookcaseDeleteBookcase = `
		DELETE
		FROM
			bookcase
		WHERE
			bookcase_id = ?
	`
)
