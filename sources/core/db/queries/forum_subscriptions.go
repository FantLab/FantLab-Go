package queries

const (
	ForumTopicSubscriptionDelete = "DELETE FROM f_topics_subscribers WHERE topic_id = ? AND user_id = ?"
	ForumTopicNewMessagesDelete  = "DELETE FROM f_new_messages WHERE topic_id = ? AND user_id = ?"
	ForumNewMessagesUpdate       = `
		UPDATE users
		SET new_forum_answers = (SELECT COUNT(*) FROM f_new_messages WHERE user_id = ?)
		WHERE user_id = ?
	`
	ForumTopicSubscriptionInsert = `
		INSERT INTO f_topics_subscribers (user_id, topic_id, date_of_add) VALUES (?, ?, NOW())
		ON DUPLICATE KEY UPDATE date_of_add = NOW()
	`
)
