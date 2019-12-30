package queries

const (
	AvailableForums = `
		SELECT
			g.access_to_forums
		FROM
			user_groups g
		JOIN
			users u ON u.user_group_id = g.user_group_id
		WHERE
			u.user_id = ?
	`

	Forums = `
		SELECT
			f.forum_id,
			f.name,
			f.description,
			f.topic_count,
			f.message_count,
			f.last_topic_id,
			f.last_topic_name,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			f.last_message_id,
			m.message_text AS last_message_text,
			f.last_message_date,
			fb.forum_block_id,
			fb.name AS forum_block_name
		FROM
			f_forums f
		JOIN
			f_forum_blocks fb ON fb.forum_block_id = f.forum_block_id
		LEFT JOIN
			users u ON u.user_id = f.last_user_id
		JOIN
			f_messages_text m ON m.message_id = f.last_message_id
		WHERE
			f.forum_id IN (?)
		ORDER BY
			fb.level, f.level
	`

	ForumModerators = `
		SELECT
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			md.forum_id
		FROM
			f_moderators md
		LEFT JOIN
			users u ON u.user_id = md.user_id
		ORDER BY
			md.forum_id, u.user_class DESC, u.level DESC
	`

	ForumExists = "SELECT 1 FROM f_forums WHERE forum_id = ? AND forum_id IN (?)"

	ForumTopics = `
		SELECT
			t.topic_id,
			t.name,
			t.date_of_add,
			t.views,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			t.topic_type_id,
			t.is_closed,
			t.is_pinned,
			t.message_count,
			t.last_message_id,
			u2.user_id AS last_user_id,
			u2.photo_number AS last_photo_number,
			m.message_text AS last_message_text,
			t.last_message_date
		FROM
			f_topics t
		LEFT JOIN
			users u ON u.user_id = t.user_id
		LEFT JOIN
			users u2 ON u2.user_id = t.last_user_id
		JOIN
			f_messages_text m ON m.message_id = t.last_message_id
		WHERE
			t.forum_id = ?
		ORDER BY
			t.is_pinned DESC, t.last_message_date DESC
		LIMIT ?
		OFFSET ?
	`

	ForumTopic = `
		SELECT
			t.topic_id,
			t.name,
			t.date_of_add,
			t.views,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			t.topic_type_id,
			t.is_closed,
			t.is_pinned,
			t.message_count,
			t.last_message_id,
			u2.user_id AS last_user_id,
			u2.photo_number AS last_photo_number,
			m.message_text AS last_message_text,
			t.last_message_date
		FROM
			f_topics t
		LEFT JOIN
			users u ON u.user_id = t.user_id
		LEFT JOIN
			users u2 ON u2.user_id = t.last_user_id
		JOIN
			f_forums f ON f.forum_id = t.forum_id
		JOIN
			f_messages_text m ON m.message_id = t.last_message_id
		WHERE
			t.topic_id = ? AND f.forum_id IN (?)
	`

	ForumTopicsCount = "SELECT COUNT(*) FROM f_topics WHERE forum_id = ?"

	ShortForumTopic = `
		SELECT
			t.topic_id,
			t.name AS topic_name,
			f.forum_id,
			f.name AS forum_name
		FROM
			f_topics t
		JOIN
			f_forums f ON f.forum_id = t.forum_id
		WHERE
			t.topic_id = ? AND t.forum_id IN (?)
	`

	ForumTopicMessagesCount = "SELECT COUNT(*) FROM f_messages WHERE topic_id = ?"

	// todo не нужны ли какие-нибудь доп. манипуляции с полем number при чтении
	//  (например, при переносе сообщений между темами)?
	//  https://github.com/parserpro/fantlab/blob/HEAD@%7B2019-06-17T18:16:10Z%7D/pm/Forum.pm#L1011
	ForumTopicMessages = `
		SELECT
			f.message_id,
			f.date_of_add,
			f.user_id,
			u.login,
			u.sex,
			u.photo_number,
			u.user_class,
			u.sign,
			m.message_text,
			f.is_censored,
			f.vote_plus,
			ABS(f.vote_minus) as vote_minus
		FROM
			f_messages f
		LEFT JOIN
			users u ON u.user_id = f.user_id
		JOIN
			f_messages_text m ON m.message_id = f.message_id
		WHERE
			f.topic_id = ? AND f.number >= ? AND f.number <= ?
		ORDER BY
			f.date_of_add %s
	`
)
