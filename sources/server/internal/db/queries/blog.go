package queries

const (
	Communities = `
		SELECT
			b.blog_id,
			b.name,
			b.description,
			b.topics_count,
			b.is_public,
			b.last_topic_date,
			b.last_topic_head,
			b.last_topic_id,
			b.subscriber_count,
			u.user_id AS last_user_id,
			u.login AS last_login,
			u.sex AS last_sex,
			u.photo_number AS last_photo_number
		FROM
			b_blogs b
		LEFT JOIN
			users u ON u.user_id = b.last_user_id
		WHERE
			b.is_community = 1 AND b.is_hidden = 0
		ORDER BY
			b.is_public DESC, b.last_topic_date DESC
	`

	Community = `
		SELECT
			b.name,
			b.description,
			b.topics_count,
			b.is_public,
			b.last_topic_date,
			b.last_topic_head,
			b.last_topic_id,
			b.subscriber_count,
			u.user_id AS last_user_id,
			u.login AS last_login,
			u.sex AS last_sex,
			u.photo_number AS last_photo_number
		FROM
			b_blogs b
		LEFT JOIN
			users u ON u.user_id = b.last_user_id
		WHERE
			b.blog_id = ? AND b.is_community = 1 AND b.is_hidden = 0
	`

	ShortCommunity = `
		SELECT
			blog_id,
			name,
			rules
		FROM
			b_blogs
		WHERE
			blog_id = ? AND is_community = 1
	`

	CommunityModerators = `
		SELECT
			cm.user_id,
			u.login,
			u.sex,
			u.photo_number
		FROM
			b_community_moderators cm
		LEFT JOIN
			users u ON u.user_id = cm.user_id
		WHERE
			cm.blog_id = ?
		ORDER BY
			cm.comm_moder_id
	`

	CommunityAuthors = `
		SELECT
			cu.user_id,
			u.login,
			u.sex,
			u.photo_number
		FROM
			b_community_users cu
		LEFT JOIN
			users u ON u.user_id = cu.user_id
		WHERE
			cu.blog_id = ? AND cu.accepted = 1
		ORDER BY
			cu.community_user_id
	`

	CommunityTopics = `
		SELECT
			b.topic_id,
			b.head_topic,
			b.date_of_add,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			t.message_text,
			b.tags,
			b.likes_count,
			b.comments_count
		FROM
			b_topics b
		JOIN
			b_topics_text t ON t.message_id = b.topic_id
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.blog_id = ? AND b.is_opened = 1
		ORDER BY
			b.date_of_add DESC
	`

	CommunityTopicCount = `
		SELECT
			COUNT(*)
		FROM
			b_topics
		WHERE
			blog_id = ? AND is_opened = 1
	`

	Blogs = `
		SELECT
			b.blog_id,
			u.user_id,
			u.login,
			u.fio,
			u.sex,
			u.photo_number,
			b.topics_count,
			b.subscriber_count,
			b.is_close,
			b.last_topic_date,
			b.last_topic_head,
			b.last_topic_id
		FROM
			b_blogs b
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.is_community = 0 AND b.topics_count > 0
		ORDER BY
			b.is_close, b.%s DESC
		LIMIT ?
		OFFSET ?
	`

	BlogCount = `
		SELECT
			COUNT(*)
		FROM
			b_blogs
		WHERE
			is_community = 0 AND topics_count > 0
	`

	Blog = `
		SELECT
			b.blog_id,
			u.user_id,
			u.login,
			u.fio,
			u.sex,
			u.photo_number,
			b.topics_count,
			b.subscriber_count,
			b.is_close,
			b.last_topic_date,
			b.last_topic_head,
			b.last_topic_id
		FROM
			b_blogs b
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.blog_id = ? AND b.is_community = 0
	`

	BlogExists = "SELECT 1 FROM b_blogs WHERE blog_id = ? AND is_community = 0"

	BlogTopics = `
		SELECT
			b.topic_id,
			b.head_topic,
			b.date_of_add,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			t.message_text,
			b.tags,
			b.likes_count,
			b.comments_count
		FROM
			b_topics b
		JOIN
			b_topics_text t ON t.message_id = b.topic_id
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.blog_id = ? AND b.is_opened = 1
		ORDER BY
			b.date_of_add DESC
	`

	BlogTopicCount = `
		SELECT
			COUNT(*)
		FROM
			b_topics
		WHERE
			blog_id = ? AND is_opened = 1
	`

	BlogTopic = `
		SELECT
			b.topic_id,
			b.head_topic,
			b.date_of_add,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			t.message_text,
			b.tags,
			b.likes_count,
			b.comments_count
		FROM
			b_topics b
		JOIN
			b_topics_text t ON t.message_id = b.topic_id
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.topic_id = ? AND b.is_opened > 0
	`

	BlogArticleComments = `
		WITH RECURSIVE CTE (message_id, parent_message_id, is_censored, date_of_add, user_id, depth) AS (
			(
				SELECT message_id, parent_message_id, is_censored, date_of_add, user_id, 0 AS depth
				FROM b_messages
				WHERE topic_id = ? AND parent_message_id = 0 AND date_of_add > ?
				ORDER BY date_of_add %s
				LIMIT ?
			)
			UNION ALL
			SELECT bm.message_id, bm.parent_message_id, bm.is_censored, bm.date_of_add, bm.user_id, cte.depth + 1 AS depth
			FROM b_messages bm
			INNER JOIN cte ON bm.parent_message_id = cte.message_id
			WHERE depth + 1 < ?
		)
		SELECT c.message_id, c.parent_message_id, c.date_of_add, c.is_censored, u.user_id, u.login, u.sex, u.photo_number, IF(c.is_censored, '', bmt.message_text) AS content
		FROM cte c
		LEFT JOIN users u ON c.user_id = u.user_id
		LEFT JOIN b_messages_text bmt ON bmt.message_id = c.message_id
	`

	BlogArticleCommentsCount = "SELECT COUNT(*) FROM b_messages WHERE topic_id = ?"
)
