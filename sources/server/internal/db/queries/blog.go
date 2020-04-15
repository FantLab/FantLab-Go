package queries

const (
	NewBlogTopicMessagesTable = "b_new_messages"
)

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
			b.blog_id = ?
	`

	BlogExists = "SELECT 1 FROM b_blogs WHERE blog_id = ?"

	BlogGetRelatedBlogs = `
		SELECT
			blog_id
		FROM
			b_topic_blog_links
		WHERE
			topic_id = ?
	`

	BlogIsUserReadOnly = `
		SELECT
			1
		FROM
			b_readonly
		WHERE
			user_id = ? AND (blog_id = ? OR blog_id IN (?))
	`

	BlogTopics = `
		SELECT
			b.topic_id,
			b.head_topic,
			b.date_of_add,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			b.tags,
			b.likes_count,
			b.comments_count
		FROM
			b_topics b
		LEFT JOIN
			users u ON u.user_id = b.user_id
		WHERE
			b.blog_id = ? AND b.is_opened = 1
		ORDER BY
			b.date_of_add DESC
		LIMIT ?
		OFFSET ?
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
			b.blog_id,
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

	BlogTopicMessages = `
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

	BlogGetTopicMessagesCount = `
		SELECT 
			COUNT(*) 
		FROM 
			b_messages 
		WHERE 
			topic_id = ?
	`

	BlogGetTopicMessage = `
		SELECT 
			m.message_id, 
			m.topic_id,
			m.parent_message_id, 
			m.date_of_add,
			m.is_censored, 
			u.user_id, 
			u.login, 
			u.sex, 
			u.photo_number, 
			IF(m.is_censored, '', mt.message_text) AS content
		FROM 
			b_messages m
		LEFT JOIN 
			users u ON m.user_id = u.user_id
		LEFT JOIN 
			b_messages_text mt ON mt.message_id = m.message_id
		WHERE
			m.message_id = ?
	`

	BlogTopicInsertNewMessage = `
		INSERT INTO
			b_messages (
				topic_id,
				user_id,
				parent_message_id,
				message_length,
				is_censored,
				date_of_add,
				topic_type
			)
		VALUES
			(?, ?, ?, ?, ?, ?, 0)
	`

	BlogSetMessageText = `
		REPLACE
			b_messages_text
		SET
			message_id = ?,
			message_text = ?
	`

	BlogUpdateLastCommentReadActuality = `
		UPDATE
			b_last_comment_read
		SET
			is_actual = 0
		WHERE
			topic_id = ?
	`

	BlogUpdateTopicCommentCount = `
		UPDATE
			b_topics
		SET
			comments_count = ?
		WHERE
			topic_id = ?
	`

	BlogIncrementNewBlogCommentsCount = `
		UPDATE
			users
		SET
			new_blog_comments = new_blog_comments + 1
		WHERE
			user_id IN (?)
	`

	BlogGetTopicSubscribers = `
		SELECT
			user_id
		FROM
			b_topics_subscribers
		WHERE
			topic_id = ? AND user_id != ?
	`

	BlogGetFirstLevelMessageCount = `
		SELECT
			COUNT(*)
		FROM
			b_messages
		WHERE
			topic_id = ? AND topic_type = 0 AND parent_message_id = 0 AND message_id <= ?
	`

	BlogGetUserIsCommunityModerator = `
		SELECT
			COUNT(*)
		FROM
			b_community_moderators
		WHERE
			blog_id = ? AND user_id = ?
	`

	BlogGetUserIsCommunityTopicModerator = `
		SELECT
			COUNT(*)
		FROM
			b_topic_blog_links b
		LEFT JOIN
			b_community_moderators bm ON bm.blog_id = b.blog_id
		WHERE
			b.topic_id = ? AND bm.user_id = ?
	`

	BlogDeleteMessage = `
		DELETE FROM
			b_messages
		WHERE
			message_id = ?
	`

	BlogDeleteMessageText = `
		DELETE FROM
			b_topics_text
		WHERE
			message_id = ?
	`

	BlogUpdateMessagesParent = `
		UPDATE
			b_messages
		SET
			parent_message_id = ?
		WHERE
			parent_message_id = ?
	`

	BlogDeleteNewMessage = `
		DELETE FROM
			b_new_messages
		WHERE
			message_id = ?
	`

	BlogInsertMessageDeleted = `
		REPLACE INTO
			b_message_delete_topics
		SET
			topic_id = ?
	`
)
