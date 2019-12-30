package queries

const (
	WorkExists   = "SELECT 1 FROM works WHERE work_id = ?"
	WorkUserMark = "SELECT mark FROM marks2 WHERE user_id = ? AND work_id = ?"
	WorkChildren = `
		WITH RECURSIVE CTE (work_id, parent_work_id, is_bonus) AS (
			SELECT work_id, parent_work_id, is_bonus
			FROM work_links
			WHERE parent_work_id = ?
			UNION ALL
			SELECT wl.work_id, wl.parent_work_id, wl.is_bonus
			FROM work_links wl
			INNER JOIN cte ON wl.parent_work_id = cte.work_id
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
	`
)
