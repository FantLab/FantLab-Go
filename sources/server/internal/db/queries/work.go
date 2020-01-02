package queries

const (
	WorkExists   = "SELECT 1 FROM works WHERE work_id = ?"
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
)
