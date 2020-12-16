package queries

const (
	ForumTopicSubscriptionDelete = `
		DELETE
		FROM
			f_topics_subscribers
		WHERE
			topic_id = ? AND user_id = ?
	`

	ForumTopicNewMessagesDelete = `
		DELETE
		FROM
			f_new_messages
		WHERE
			topic_id = ? AND user_id = ?
	`

	ForumNewMessagesUpdate = `
		UPDATE
			users
		SET
			new_forum_answers = (SELECT COUNT(*) FROM f_new_messages WHERE user_id = ?)
		WHERE 
			user_id = ?
	`

	ForumTopicSubscriptionInsert = `
		INSERT INTO 
			f_topics_subscribers (
				user_id, 
				topic_id, 
				date_of_add
			)
		VALUES
			(?, ?, NOW())
		ON DUPLICATE KEY UPDATE
			date_of_add = NOW()
	`

	// Traditionally, an EXISTS subquery starts with SELECT *, but it could begin with SELECT 5 or SELECT column1 or anything at all.
	// MySQL ignores the SELECT list in such a subquery, so it makes no difference.
	// (https://dev.mysql.com/doc/refman/8.0/en/exists-and-not-exists-subqueries.html)
	ForumGetTopicSubscriptionExists = `
		SELECT
			EXISTS (
				SELECT
					*
				FROM
					f_topics_subscribers
				WHERE
					topic_id = ? AND user_id = ?
			)
	`
)
