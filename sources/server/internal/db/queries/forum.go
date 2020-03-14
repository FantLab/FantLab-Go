package queries

const (
	NewForumAnswersTable = "f_new_messages"
)

const (
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

	ForumGetShortForum = `
		SELECT
			forum_id,
			only_for_admins
		FROM
			f_forums
		WHERE
			forum_id = ?
	`

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
			t.last_message_date,
			t.moderated
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
			t.forum_id,
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
			t.last_message_date,
			t.moderated
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
			t.is_firstpost AS is_first_message_pinned,
			f.forum_id,
			f.name AS forum_name
		FROM
			f_topics t
		JOIN
			f_forums f ON f.forum_id = t.forum_id
		WHERE
			t.topic_id = ? AND t.forum_id IN (?)
	`

	ForumTopicGetIsEditTopicStarter = `
		SELECT
			t.is_edit_topicstarter
		FROM
			f_topics t
		JOIN
			f_messages m ON m.topic_id = t.topic_id
		WHERE
			m.message_id = ?
	`

	ForumTopicMessagesCount = "SELECT COUNT(*) FROM f_messages WHERE topic_id = ?"

	// TODO Не нужны ли какие-нибудь доп. манипуляции с полем number при чтении
	//  (например, при переносе сообщений между темами)?
	//  https://github.com/parserpro/fantlab/blob/HEAD@%7B2019-06-17T18:16:10Z%7D/pm/Forum.pm#L1011
	ForumTopicMessages = `
		SELECT
			f.message_id,
			f.topic_id,
			f.date_of_add,
			f.user_id,
			f.is_red,
			u.login,
			u.sex,
			u.photo_number,
			u.user_class,
			u.sign,
			m.message_text,
			f.is_censored,
			f.vote_plus,
			ABS(f.vote_minus) AS vote_minus,
			f.number
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

	ForumTopicFirstMessage = `
		SELECT
			f.message_id,
			f.topic_id,
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
			ABS(f.vote_minus) AS vote_minus,
			f.number
		FROM
			f_messages f
		LEFT JOIN
			users u ON u.user_id = f.user_id
		JOIN
			f_messages_text m ON m.message_id = f.message_id
		WHERE
			f.topic_id = ? AND f.number = 1
	`

	ForumGetShortMessage = `
		SELECT
			message_id,
			topic_id,
			forum_id,
			user_id,
			is_censored,
			is_red,
			date_of_add,
			vote_plus,
			vote_minus,
			number
		FROM
			f_messages
		WHERE
			message_id = ? AND forum_id IN (?)
	`

	UserIsForumModerator = `
		SELECT
			COUNT(*)
		FROM
			f_topics ft 
		INNER JOIN 
			f_moderators fmd ON ft.forum_id = fmd.forum_id 
		WHERE
			fmd.user_id = ? AND ft.topic_id = ?
	`

	ForumInsertNewMessage = `
		INSERT INTO
			f_messages (
				message_length,
				topic_id,
				user_id,
				forum_id,
				is_red,
				date_of_add,
				date_of_edit,
				number
			)
		SELECT
			?, ?, ?, ?, ?, NOW(), NOW(), COALESCE(MAX(number), 0) + 1
		FROM
			f_messages
		WHERE
			topic_id = ?
	`

	ForumGetTopicLastMessage = `
		SELECT 
			* 
		FROM 
			f_messages 
		WHERE
			topic_id = ?
		ORDER BY 
			message_id DESC 
		LIMIT 1
	`

	ForumSetMessageText = `
		REPLACE
			f_messages_text
		SET
			message_id = ?,
			message_text = ?
	`

	ForumCancelTopicMessagePreview = `
		DELETE FROM
			f_messages_preview
		WHERE
			user_id = ? AND topic_id = ?
	`

	ForumUpdateUserStat = `
		UPDATE
			users
		SET
			messagecount = messagecount + 1,
			need_recalc_level = 1
		WHERE
			user_id = ?
	`

	// need_update_numbers - требует пересчета number для сообщений
	ForumSetTopicLastMessage = `
		UPDATE
			f_topics
		SET
			message_count = message_count + 1,
			last_message_id = ?,
			last_user_id = ?,
			last_user_name = ?,
			last_message_date = NOW(),
			need_update_numbers = 1
		WHERE
			topic_id = ?
	`

	// need_sindex - тема требует переиндексации Sphinx-ом
	ForumMarkTopicUpdated = `
		UPDATE
			f_topics
		SET
			need_sindex = 1
		WHERE
			topic_id = ?
	`

	// SUM возвращает значение типа DECIMAL (https://dev.mysql.com/doc/refman/8.0/en/group-by-functions.html)
	ForumGetStat = `
		SELECT
			COUNT(*) AS topic_count,
			CAST(SUM(message_count) AS SIGNED) AS forum_message_count
		FROM
			f_topics
		WHERE
			forum_id = ? AND moderated = 1
	`

	ForumGetTopicMessageCount = `
		SELECT
			message_count
		FROM
			f_topics
		WHERE
			topic_id = ?
	`

	ForumSetForumLastTopic = `
		UPDATE
			f_forums
		SET
			message_count = ?,
			topic_count = ?,
			last_message_id = ?,
			last_user_id = ?,
			last_user_name = ?,
			last_topic_id = ?,
			last_topic_name = ?,
			last_message_date = NOW(),
			last_topic_page_count = ?
		WHERE
			forum_id = ?
	`

	ForumGetTopicSubscribers = `
		SELECT
			user_id
		FROM
			f_topics_subscribers
		WHERE
			topic_id = ? AND user_id != ?
	`

	ForumUpdateNewForumAnswersCount = `
		UPDATE
			users
		SET
			new_forum_answers = new_forum_answers + 1
		WHERE
			user_id IN (?)
	`

	ForumUpdateMessage = `
		UPDATE
			f_messages
		SET
			message_length = ?,
			is_red = ?
		WHERE
			message_id = ?
	`
)
