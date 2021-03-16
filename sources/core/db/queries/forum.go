package queries

const (
	NewForumAnswersTable = "f_new_messages"
)

const (
	ForumGetForums = `
		SELECT
			f.forum_id,
			f.name,
			f.description,
			f.topic_count,
			f.message_count,
			f.last_topic_id,
			f.last_topic_name,
			f.last_message_id,
			u.user_id AS last_message_user_id,
			u.login AS last_message_user_login,
			u.sex AS last_message_user_sex,
			u.photo_number AS last_message_user_photo_number,
			mt.message_text AS last_message_text,
			f.last_message_date,
			fb.forum_block_id,
			fb.name AS forum_block_name,
			f.not_moderated_topic_count
		FROM
			f_forums f
		JOIN
			f_forum_blocks fb ON fb.forum_block_id = f.forum_block_id
		LEFT JOIN
			users u ON u.user_id = f.last_user_id
		LEFT JOIN
			f_messages_text mt ON mt.message_id = f.last_message_id
		WHERE
			f.forum_id IN (?)
		ORDER BY
			fb.level, f.level
	`

	ForumGetModerators = `
		SELECT
			md.forum_id,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number
		FROM
			f_moderators md
		LEFT JOIN
			users u ON u.user_id = md.user_id
		ORDER BY
			md.forum_id, u.user_class DESC, u.level DESC
	`

	ForumGetNotModeratedTopicIds = `
		SELECT
			topic_id
		FROM
			f_topics
		WHERE
			moderated = 0
	`

	// SUM возвращает значение типа DECIMAL (https://dev.mysql.com/doc/refman/8.0/en/group-by-functions.html)
	ForumGetNotReadMessageCounts = `
		SELECT
			f.forum_id,
			f.message_count - CAST(SUM(utr.read_count) AS SIGNED) AS not_read_message_count
		FROM
			user_topic_reads utr
		JOIN
			f_forums f ON f.forum_id = utr.forum_id
		WHERE
			utr.user_id = ? AND utr.topic_id NOT IN (?)
		GROUP BY
			f.forum_id
	`

	ForumGetForumExists = `
		SELECT
			EXISTS (
				SELECT
					*
				FROM
					f_forums
				WHERE
					forum_id = ? AND forum_id IN (?)
			)
	`

	ForumGetForum = `
		SELECT
			forum_id,
			name,
			only_for_admins,
			forum_closed
		FROM
			f_forums
		WHERE
			forum_id = ?
	`

	ForumGetForumModerators = `
		SELECT
			u.user_id,
			u.login,
			u.sex,
			u.photo_number
		FROM
			f_moderators md
		LEFT JOIN
			users u ON u.user_id = md.user_id
		WHERE
			md.forum_id = ?
		ORDER BY
			u.user_class DESC, u.level DESC
	`

	// NOTE: странное условие в запросе - попытка убить сразу двух зайцев:
	// 1. не заниматься формированием запроса в рантайме (в отличие от Perl-бэка)
	// 2. обработать сразу несколько кейсов, связанных с видимостью тем:
	//   a. незалогиненный пользователь - должны быть видны все отмодерированные темы
	//   b. обычный пользователь - должны быть видны все отмодерированные темы и те, автор которых - он сам
	//   c. модератор этого форума - должны быть видны все темы, в том числе неотмодерированные
	ForumGetForumTopics = `
		SELECT
			t.topic_id,
			t.name,
			t.date_of_add,
			t.views,
			u.user_id,
			u.login AS user_login,
			u.sex AS user_sex,
			u.photo_number AS user_photo_number,
			t.topic_type_id,
			t.is_closed,
			t.is_pinned,
			t.message_count,
			t.last_message_id,
			u2.user_id AS last_message_user_id,
			u2.photo_number AS last_message_user_photo_number,
			u2.login AS last_message_user_login,
			u2.sex AS last_message_user_sex,
			m.message_text AS last_message_text,
			t.last_message_date,
			t.moderated
		FROM
			f_topics t
		LEFT JOIN
			users u ON u.user_id = t.user_id
		LEFT JOIN
			users u2 ON u2.user_id = t.last_user_id
		LEFT JOIN
			f_messages_text m ON m.message_id = t.last_message_id
		WHERE
			t.forum_id = ? AND (t.moderated IN (?) OR t.user_id = ?)
		ORDER BY
			t.is_pinned DESC, t.last_message_date DESC
		LIMIT ?
		OFFSET ?
	`

	ForumGetForumTopicCount = `
		SELECT
			COUNT(*)
		FROM
			f_topics
		WHERE
			forum_id = ? AND (moderated IN (?) OR user_id = ?)
	`

	ForumGetTopicsSubscriptions = `
		SELECT
			topic_id
		FROM
			f_topics_subscribers
		WHERE
			user_id = ? AND topic_id IN (?)
	`

	ForumGetTopicsDatesOfRead = `
		SELECT
			topic_id,
			date_of_read
		FROM
			user_topic_reads
		WHERE
			user_id = ? AND topic_id IN (?)
	`

	ForumGetNotReadTopicInfo = `
		SELECT
			topic_id,
			MIN(message_id) AS first_not_read_message_id,
			COUNT(*) AS not_read_message_count
		FROM
			f_messages
		WHERE
			topic_id = ? AND date_of_add > ?
		GROUP BY
			topic_id
	`

	ForumGetTopicExists = `
		SELECT
			EXISTS (
				SELECT
					*
				FROM
					f_topics
				WHERE
					topic_id = ? AND forum_id IN (?)
			)
	`

	ForumGetTopic = `
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
			u2.login AS last_login,
			u2.sex AS last_sex,
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
		LEFT JOIN
			f_messages_text m ON m.message_id = t.last_message_id
		WHERE
			t.topic_id = ?
	`

	ForumGetTopicShort = `
		SELECT
			t.topic_id,
			t.name AS topic_name,
			t.is_firstpost AS is_first_message_pinned,
			t.topic_type_id,
			t.is_closed,
			t.is_edit_topicstarter,
			f.forum_id,
			f.name AS forum_name
		FROM
			f_topics t
		JOIN
			f_forums f ON f.forum_id = t.forum_id
		WHERE
			t.topic_id = ?
	`

	ForumTopicMessageCount = `
		SELECT
			COUNT(*)
		FROM
			f_messages
		WHERE
			topic_id = ?
	`

	// NOTE Этот запрос не всегда отдает корректные результаты. Дело в том, что мы опираемся на поле number, а оно может
	// быть непоследовательным - например, после удаления сообщений или переноса их между темами это число для соседних
	// по времени сообщений может отличаться в разы. В Perl-бэке исправлением number занимает Cron-скрипт
	// update_forum_and_pm_numbers.pm. Пока он не отработает, корректость результатов не гарантирована.
	ForumGetTopicMessageIds = `
		SELECT
			message_id
		FROM
			f_messages
		WHERE
			topic_id = ? AND number >= ? AND number <= ?
	`

	ForumGetTopicMessages = `
		SELECT
			f.message_id,
			f.topic_id,
			f.forum_id,
			f.date_of_add,
			f.is_red,
			f.is_censored,
			f.vote_plus,
			ABS(f.vote_minus) AS vote_minus,
			f.number,
			m.message_text,
			u.user_id,
			u.login,
			u.sex,
			u.photo_number,
			u.user_class,
			u.sign,
			u.approved
		FROM
			f_messages f
		LEFT JOIN
			users u ON u.user_id = f.user_id
		LEFT JOIN
			f_messages_text m ON m.message_id = f.message_id
		WHERE
			f.message_id IN (?)
		ORDER BY
			f.date_of_add %s
	`

	ForumGetTopicMessagesAttachments = `
		SELECT
			message_id,
			file_name,
			file_size
		FROM
			f_files
		WHERE
			file_group = 'forum' AND message_id IN (?)
	`

	ForumGetTopicFirstMessageId = `
		SELECT
			f.message_id
		FROM
			f_messages f
		WHERE
			f.topic_id = ? AND f.number = 1
	`

	ForumGetMessageExists = `
		SELECT
			EXISTS (
				SELECT
					*
				FROM
					f_messages
				WHERE
					message_id = ? AND forum_id IN (?)
			)
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
			message_id = ?
	`

	ForumGetTopicReadDate = `
		SELECT
			date_of_read
		FROM
			user_topic_reads
		WHERE
			topic_id = ? AND user_id = ?
	`

	ForumGetVotedMessageIds = `
		SELECT
			message_id
		FROM
			f_messages_votes
		WHERE
			message_id IN (?) AND user_id = ?
	`

	ForumGetWarnedMessageIds = `
		SELECT
			msg_id
		FROM
			user_fine
		WHERE
			msg_id IN (?)
	`

	ForumGetModerCalledMessageIds = `
		SELECT
			msg_id
		FROM
			f_callmoder
		WHERE
			msg_id IN (?)
	`

	ForumGetTopicAnswers = `
		SELECT
			number,
			name,
			choices
		FROM
			f_answers
		WHERE
			topic_id = ?
		ORDER BY
			number ASC
	`

	ForumGetUserTopicAnswerExists = `
		SELECT
			EXISTS (
				SELECT
					*
				FROM
					f_user_topic_answer
				WHERE
					topic_id = ? AND user_id = ?
			)
	`

	ForumGetTopicAnsweredUsers = `
		SELECT
			u.user_id,
			u.login
		FROM
			f_user_topic_answer futa
		LEFT JOIN
			users u ON u.user_id = futa.user_id
		WHERE
			futa.topic_id = ?
		ORDER BY
			futa.date_of_add ASC
	`

	ForumGetUserIsForumModerator = `
		SELECT
			EXISTS (
				SELECT
					*
				FROM
					f_moderators
				WHERE
					user_id = ? AND forum_id = ?
			)
	`

	ForumIncrementTopicViewCount = `
		UPDATE
			f_topics
		SET
			views = views + 1
		WHERE
			topic_id = ?
	`

	ForumDeleteUserForumNewMessages = `
		DELETE
		FROM
			f_new_messages
		WHERE
			user_id = ? AND topic_id = ? AND (message_id IN (?) OR date_of_add <= ?)
	`

	ForumGetUserTopicReadCount = `
		SELECT
			read_count
		FROM
			user_topic_reads
		WHERE
			user_id = ? AND topic_id = ?
	`

	ForumInsertUserTopicReadDate = `
		INSERT INTO
			user_topic_reads (
				user_id,
				topic_id,
				date_of_read,
				read_count,
				forum_id
			)
		VALUES
			(?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			date_of_read = ?,
			read_count = ?,
			forum_id = ?
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

	ForumSetMessageText = `
		REPLACE
			f_messages_text
		SET
			message_id = ?,
			message_text = ?
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

	ForumSetTopicLastMessage = `
		UPDATE
			f_topics t
		SET
			t.message_count = (SELECT COUNT(DISTINCT m.message_id) FROM f_messages m WHERE m.topic_id = t.topic_id),
			t.last_message_id = ?,
			t.last_user_id = ?,
			t.last_user_name = ?,
			t.last_message_date = ?,
			t.need_update_numbers = 1
		WHERE
			t.topic_id = ?
	`

	// need_sindex - тема требует переиндексации Sphinx-ом
	ForumMarkTopicNeedSphinxReindex = `
		UPDATE
			f_topics
		SET
			need_sindex = 1
		WHERE
			topic_id = ?
	`

	// SUM возвращает значение типа DECIMAL (https://dev.mysql.com/doc/refman/8.0/en/group-by-functions.html)
	ForumGetForumStat = `
		SELECT
			COUNT(*) AS topic_count,
			CAST(SUM(message_count) AS SIGNED) AS message_count
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
			last_message_date = ?,
			last_topic_page_count = ?,
			not_moderated_topic_count = ?
		WHERE
			forum_id = ?
	`

	ForumGetTopicSubscribers = `
		SELECT
			user_id
		FROM
			f_topics_subscribers
		WHERE
			topic_id = ?
	`

	ForumIncrementNewForumAnswersCount = `
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

	ForumDeleteMessage = `
		DELETE
		FROM
			f_messages
		WHERE
			message_id = ?
	`

	ForumDeleteMessageText = `
		DELETE
		FROM
			f_messages_text
		WHERE
			message_id = ?
	`

	ForumDeleteMessageFiles = `
		DELETE
		FROM
			f_files
		WHERE
			file_group = 'forum' AND message_id = ?
	`

	ForumMarkMessageDeleted = `
		INSERT INTO
			f_messages_deleted (message_id)
		VALUES
			(?)
	`

	ForumMarkTopicNeedRecalc = `
		UPDATE
			f_topics
		SET
			need_recalc_unread_count = 1
		WHERE
			topic_id = ?
	`

	ForumUpdateUserTopicReads = `
		UPDATE
			user_topic_reads
		SET
			read_count = read_count - 1
		WHERE
			topic_id = ? AND date_of_read >= ?
	`

	ForumGetTopicStat = `
		SELECT
			MAX(message_id) AS last_message_id
		FROM
			f_messages
		WHERE
			topic_id = ?
	`

	ForumGetMessageInfo = `
		SELECT
			m.user_id,
			u.login,
			m.date_of_add
		FROM
			f_messages m
		LEFT JOIN
			users u ON u.user_id = m.user_id
		WHERE
			m.message_id = ?
	`

	ForumGetNotModeratedTopicCount = `
		SELECT
			COUNT(*) AS topic_count
		FROM
			f_topics
		WHERE
			forum_id = ? AND moderated = 0
	`

	ForumGetLastTopic = `
		SELECT
			last_message_id,
			topic_id,
			name,
			last_user_id,
			last_user_name AS last_login,
			last_message_date,
			message_count
		FROM
			f_topics
		WHERE
			forum_id = ? AND moderated = 1
		ORDER BY
			last_message_date DESC
		LIMIT 1
	`

	ForumDeleteNewForumAnswer = `
		DELETE
		FROM
			f_new_messages
		WHERE
			message_id = ?
	`

	ForumDecrementNewForumAnswersCount = `
		UPDATE
			users
		SET
			new_forum_answers = new_forum_answers - 1
		WHERE
			user_id IN (?) AND new_forum_answers > 0
	`
)
