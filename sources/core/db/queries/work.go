package queries

const (
	WorkExists = "SELECT 1 FROM works WHERE work_id = ?"

	WorkGetWork = `
		SELECT
			work_id,
			name,
			autor_id,
			autor2_id,
			autor3_id,
			autor4_id,
			autor5_id
		FROM
			works
		WHERE
			work_id = ?
	`

	WorkGetWorks = `
		SELECT
			work_id,
			name
		FROM
			works
		WHERE
			work_id IN (?)
	`

	WorkUserMark = "SELECT mark FROM marks2 WHERE user_id = ? AND work_id = ?"
	WorkChildren = `
		WITH RECURSIVE CTE (work_id, parent_work_id, is_bonus, group_index, link, depth) AS (
			(
				SELECT work_id, parent_work_id, is_bonus, group_index, 1 AS link, 0 AS depth
				FROM work_links
				WHERE parent_work_id = ?
				UNION
				SELECT work_id, parent_work_id, is_bonus, group_index, 0 AS link, 0 AS depth
				FROM works
				WHERE parent_work_id = ?
			)
			UNION ALL
			SELECT wl.work_id, wl.parent_work_id, wl.is_bonus, wl.group_index, 1 AS link, cte.depth + 1 AS depth
			FROM work_links wl
			INNER JOIN cte ON wl.parent_work_id = cte.work_id
			WHERE depth + 1 < ?
		)
		SELECT
			c.*,
			w.name,
			w.rusname,
			w.year,
			w.show_subworks_in_biblio,
			w.not_finished,
			w.is_plan,
			w.published,
			w.work_type_id,
			ws.midmark_by_weight,
			ws.markcount,
			ws.responsecount
		FROM cte c
		LEFT JOIN works w ON w.work_id = c.work_id
		LEFT JOIN work_stats ws ON ws.work_id = w.work_id
		ORDER BY c.group_index, w.year
	`

	WorkMarkWorksNeedPopularityRecalc = `
		UPDATE
			works
		SET
			popularity_need_recalc = 1
		WHERE
			work_id IN (?)
	`

	WorkGetRegisteredWorkAutorIds = `
		SELECT
			u.autor_id
		FROM
			works w
		LEFT JOIN
			users u ON (u.autor_id = w.autor_id OR u.autor_id = w.autor2_id OR u.autor_id = w.autor3_id OR u.autor_id = w.autor4_id OR u.autor_id = w.autor5_id)
		WHERE
			w.work_id = ?
	`

	WorkInsertResponseCount = `
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
	WorkDecrementResponseCount = `
		UPDATE
			work_stats
		SET
			responsecount = responsecount - 1,
			responsecount1 = responsecount1 - 1
		WHERE
			work_id = ?
	`
)
